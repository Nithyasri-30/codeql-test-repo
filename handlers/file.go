package handlers

import (
	"net/http"
	"path/filepath"

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

		// Serve the requested file from uploads directory
		fullPath := filepath.Join(uploadsDir, filename)
		c.File(fullPath)
	}
}
