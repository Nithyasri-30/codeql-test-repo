package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"codeql-test-repo/models"

	"github.com/gin-gonic/gin"
)

func GetProfile(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var user models.User
		err = db.QueryRow(
			"SELECT id, username, email, bio, created_at FROM users WHERE id = ?",
			id,
		).Scan(&user.ID, &user.Username, &user.Email, &user.Bio, &user.CreatedAt)

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
