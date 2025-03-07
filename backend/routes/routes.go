package routes

import (
	"sers/handlers"
	"sers/middleware"
	"sers/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// API routes
	api := router.Group("/api")
	{
		// Auth routes
		authHandler := handlers.NewAuthHandler(db)
		api.POST("/auth/register", authHandler.Register)
		api.POST("/auth/login", authHandler.Login)

		sosService := services.NewSOSService(db)
		sosHandler := handlers.NewSOSHandler(sosService)
		api.POST("/sos/trigger", middleware.AuthMiddleware(), sosHandler.TriggerSOS)
	}
}
