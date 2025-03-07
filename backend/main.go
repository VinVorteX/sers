package main

import (
	"log"
	"sers/config"
	"sers/middleware"
	"sers/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize database
	db := config.InitDB()

	// Initialize router
	router := gin.Default()

	// CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Setup routes
	routes.SetupRoutes(router, db)

	// Start server
	log.Printf("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
