package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-api/db"
)

// GetTodosHandler retrieves a list of todos from the database
func GetTodosHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		todos, err := queries.ListTodos(c, userId.(int64))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
			return
		}

		c.JSON(http.StatusOK, todos)
	}
}

// GetTodoByIDHandler retrieves a single todo from the database
func GetTodoByIDHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		idString := c.Param("id")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
			return
		}

		params := db.GetTodoByIDParams{
			ID:     id,
			UserID: userId.(int64),
		}

		todo, err := queries.GetTodoByID(c, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todo"})
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}

// CreateTodoHandler creates a new todo in the database
func CreateTodoHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		var todo db.CreateTodoParams
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		params := db.CreateTodoParams{
			Title:       todo.Title,
			Description: todo.Description,
			Priority:    todo.Priority,
			UserID:      userId.(int64),
		}

		err := queries.CreateTodo(c, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully"})
	}
}

// DeleteTodoHandler deletes a todo from the database
func DeleteTodoHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		idString := c.Param("id")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
			return
		}

		params := db.DeleteTodoParams{
			ID:     id,
			UserID: userId.(int64),
		}

		err = queries.DeleteTodo(c, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
	}
}
