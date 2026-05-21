/**
 * @name Unvalidated redirect via Gin context
 * @description Detects calls to Gin's Context.Redirect() where the target URL
 *              originates from user-controlled input (query params, form data,
 *              headers) without validation. This can allow open redirect attacks
 *              where an attacker crafts a link that redirects victims to a
 *              malicious site after interacting with a trusted domain.
 * @kind path-problem
 * @problem.severity warning
 * @security-severity 6.1
 * @precision medium
 * @id go/gin-unvalidated-redirect
 * @tags security
 *       external/cwe/cwe-601
 */

import go
import semmle.go.dataflow.TaintTracking
import semmle.go.security.OpenUrlRedirectCustomizations

/**
 * A call to (*gin.Context).Redirect() as an open redirect sink.
 * Matches any method named "Redirect" on the gin.Context type,
 * treating the second argument (the URL) as the sink.
 */
class GinRedirectSink extends DataFlow::Node {
  GinRedirectSink() {
    exists(DataFlow::MethodCallNode call |
      call.getTarget().hasQualifiedName("github.com/gin-gonic/gin", "Context", "Redirect") and
      this = call.getArgument(1)
    )
  }
}

/**
 * User input sources from Gin's Context methods.
 * Covers the common ways user-controlled data enters a Gin handler:
 *   c.Query(), c.Param(), c.PostForm(), c.GetHeader(),
 *   c.DefaultQuery(), c.DefaultPostForm()
 */
class GinUserInputSource extends DataFlow::Node {
  GinUserInputSource() {
    exists(DataFlow::MethodCallNode call, string methodName |
      call
          .getTarget()
          .hasQualifiedName("github.com/gin-gonic/gin", "Context",
            ["Query", "Param", "PostForm", "GetHeader", "DefaultQuery", "DefaultPostForm"]) and
      methodName = call.getTarget().getName() and
      this = call.getResult(0)
    )
  }
}

/**
 * Taint tracking configuration: user input flowing to a Redirect call.
 */
module GinRedirectConfig implements DataFlow::ConfigSig {
  predicate isSource(DataFlow::Node source) { source instanceof GinUserInputSource }

  predicate isSink(DataFlow::Node sink) { sink instanceof GinRedirectSink }
}

module GinRedirectFlow = TaintTracking::Global<GinRedirectConfig>;

import GinRedirectFlow::PathGraph

from GinRedirectFlow::PathNode source, GinRedirectFlow::PathNode sink
where GinRedirectFlow::flowPath(source, sink)
select sink.getNode(), source, sink,
  "Unvalidated redirect: user input from $@ flows to Gin Redirect() call without validation.",
  source.getNode(), "this Gin context method"
