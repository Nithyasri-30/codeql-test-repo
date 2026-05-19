package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserProfilePage(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		var bio string
		var email string
		err := db.QueryRow(
			"SELECT email, bio FROM users WHERE username = ?", username,
		).Scan(&email, &bio)

		if err != nil {
			c.String(http.StatusNotFound, "User not found")
			return
		}

		// Render profile as HTML page
		html := fmt.Sprintf(`<html>
<head><title>%s's Profile</title></head>
<body>
  <h1>%s</h1>
  <p>Email: %s</p>
  <div class="bio">%s</div>
</body>
</html>`, username, username, email, bio)

		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	}
}
