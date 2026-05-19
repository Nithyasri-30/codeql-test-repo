package handlers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchWebhook() gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL := c.Query("url")
		if targetURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "URL parameter is required"})
			return
		}

		// Fetch the provided URL to validate webhook endpoints
		resp, err := http.Get(targetURL)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to reach URL"})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": resp.StatusCode,
			"body":   string(body),
		})
	}
}
