package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-api/internal/db"
)

// GetTodosHandler retrieves a list of todos from the database
func GetTodosHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos, err := queries.ListTodos(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
			return
		}

		c.JSON(http.StatusOK, todos)
	}
}

// CreateTodoHandler creates a new todo in the database
func CreateTodoHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo db.CreateTodoParams
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		err := queries.CreateTodo(context.Background(), todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully"})
	}
}
