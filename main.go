package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"time"
	"todo-api/internal/app"
	"todo-api/internal/db"
	"todo-api/middleware"
	"todo-api/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection pool
	pool, err := app.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer pool.Close()

	queries := db.New(pool)

	// 5 requests per second
	rateLimiter := middleware.NewRateLimiter(1*time.Second, 5)

	// Setup web server
	server := gin.Default()

	server.Use(rateLimiter.Limit())

	routes.RegisterRoutes(server, queries)

	err = server.Run(":8080")
	if err != nil {
		return
	}

}
