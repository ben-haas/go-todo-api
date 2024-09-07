package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// InitDB initializes and returns a pgx connection pool
func InitDB() (*pgxpool.Pool, error) {
	// Get the database URL from environment variables
	databaseUrl := os.Getenv("POSTGRES_URL")
	if databaseUrl == "" {
		return nil, fmt.Errorf("POSTGRES_URL environment variable not set")
	}

	// Create a new connection pool
	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	log.Println("Connected to the database successfully.")
	return pool, nil
}
