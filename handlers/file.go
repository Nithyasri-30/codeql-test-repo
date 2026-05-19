package handlers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const uploadsDir = "./uploads"

func DownloadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := c.Query("file")
		if filename == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File parameter is required"})
			return
		}

		// Sanitize: only allow the base filename, no path traversal
		cleanName := filepath.Base(filename)
		if cleanName != filename || strings.Contains(filename, "..") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filename"})
			return
		}

		fullPath := filepath.Join(uploadsDir, cleanName)

		// Verify resolved path is still within uploads directory
		absPath, err := filepath.Abs(fullPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}
		absUploads, _ := filepath.Abs(uploadsDir)
		if !strings.HasPrefix(absPath, absUploads) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			return
		}

		c.File(fullPath)
	}
}
