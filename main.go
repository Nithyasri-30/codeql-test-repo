package main

import (
	"log"

	"codeql-test-repo/database"
	"codeql-test-repo/handlers"
	"codeql-test-repo/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDB("app.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	r := gin.Default()

	// Public routes
	r.POST("/register", handlers.Register(db))
	r.POST("/login", handlers.Login(db))

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/profile/:id", handlers.GetProfile(db))
		authorized.GET("/search", handlers.SearchUsers(db))
		authorized.GET("/download", handlers.DownloadFile())
	}

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
