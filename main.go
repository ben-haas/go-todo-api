package main

import (
	"context"
	"database/sql"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"todo-api/db"
	"todo-api/middleware"
	"todo-api/routes"
)

//go:embed db/schema.sql
var ddl string

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database
	ctx := context.Background()

	sqlite, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := sqlite.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err)
	}

	queries := db.New(sqlite)

	// Setup web server
	server := gin.Default()

	server.Use(middleware.NewRateLimiter(5, 5).Limit())

	routes.RegisterRoutes(server, queries)

	err = server.Run(":8080")
	if err != nil {
		return
	}

}
