package server

import (
	"fmt"
	"log"

	"github.com/DomEscobar/erp-dev-bench/internal/config"
	"github.com/DomEscobar/erp-dev-bench/internal/handlers"
	"github.com/DomEscobar/erp-dev-bench/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config) {
	// Setup Gin
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// CORS middleware
	r.Use(cors.Default())

	// Logger middleware
	r.Use(middleware.Logger())

	// Health check
	r.GET("/health", handlers.HealthCheck)

	// API v1 routes
	api := r.Group("/api/v1")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/items", handlers.GetItems)
		api.GET("/items/:id", handlers.GetItem)
		api.POST("/items", handlers.CreateItem)
		api.PUT("/items/:id", handlers.UpdateItem)
		api.DELETE("/items/:id", handlers.DeleteItem)
	}

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s", addr)
	log.Fatal(r.Run(addr))
}
