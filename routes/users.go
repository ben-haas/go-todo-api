package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-api/internal/db"
)

// GetUsersHandler retrieves a list of users from the database
func GetUsersHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := queries.ListUsers(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

// CreateUserHandler creates a new user in the database
func CreateUserHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user db.CreateUserParams
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		err := queries.CreateUser(context.Background(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}
