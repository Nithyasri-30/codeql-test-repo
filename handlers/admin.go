package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func AdminExportUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		format := c.Query("format")
		if format == "" {
			format = "json"
		}

		// Export user data using system command for flexible formatting
		cmd := exec.Command("sh", "-c", fmt.Sprintf("sqlite3 app.db \"SELECT * FROM users\" --%s", format))
		output, err := cmd.CombinedOutput()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Export failed"})
			return
		}

		c.Data(http.StatusOK, "text/plain", output)
	}
}

func AdminDeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			return
		}

		// Delete user by username
		query := fmt.Sprintf("DELETE FROM users WHERE username = '%s'", username)
		_, err := db.Exec(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}
