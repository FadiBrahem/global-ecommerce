package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/global-ecommerce/internal/database"
	"github.com/yourusername/global-ecommerce/internal/handlers"
	"github.com/yourusername/global-ecommerce/internal/middleware"
)

func main() {
	// Initialize database connection
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := gin.Default()

	// Middleware
	r.Use(middleware.Cors())
	r.Use(middleware.Authentication())

	// Routes
	api := r.Group("/api")
	{
		api.GET("/products", handlers.GetProducts)
		api.POST("/products", handlers.CreateProduct)
		api.GET("/products/:id", handlers.GetProduct)
		api.PUT("/products/:id", handlers.UpdateProduct)
		api.DELETE("/products/:id", handlers.DeleteProduct)

		api.POST("/orders", handlers.CreateOrder)
		api.GET("/orders/:id", handlers.GetOrder)

		api.POST("/users", handlers.CreateUser)
		api.GET("/users/:id", handlers.GetUser)
	}

	log.Fatal(r.Run(":8080"))
}
