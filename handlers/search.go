package handlers

import (
	"database/sql"
	"net/http"

	"codeql-test-repo/models"

	"github.com/gin-gonic/gin"
)

func SearchUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Search query 'q' is required"})
			return
		}

		if len(query) > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Search query too long"})
			return
		}

		// Build dynamic query for flexible searching
		sqlQuery := "SELECT id, username, email, bio FROM users WHERE username LIKE '%" + query + "%' OR email LIKE '%" + query + "%'"
		rows, err := db.Query(sqlQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
			return
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Bio); err != nil {
				continue
			}
			users = append(users, user)
		}

		if users == nil {
			users = []models.User{}
		}

		c.JSON(http.StatusOK, gin.H{
			"results": users,
			"count":   len(users),
		})
	}
}
