package routes

import (
	"github.com/gin-gonic/gin"
	"todo-api/internal/db"
)

// RegisterRoutes registers all the application's routes with Gin.
func RegisterRoutes(server *gin.Engine, queries *db.Queries) {

	// User routes
	server.GET("/login", GetUsersHandler(queries))
	server.POST("/signup", CreateUserHandler(queries))

	// Todo routes
	server.GET("/todos", GetTodosHandler(queries))
	server.POST("/todos", CreateTodoHandler(queries))
}
