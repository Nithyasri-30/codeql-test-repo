package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		redirectURL := c.Query("redirect_url")
		if redirectURL == "" {
			redirectURL = "/"
		}

		// Redirect user after login
		c.Redirect(http.StatusFound, redirectURL)
	}
}
