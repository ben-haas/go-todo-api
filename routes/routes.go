package routes

import (
	"github.com/gin-gonic/gin"
	"todo-api/internal/db"
	"todo-api/middleware"
)

// RegisterRoutes registers all the application's routes with Gin.
func RegisterRoutes(server *gin.Engine, queries *db.Queries) {

	server.GET("/users", GetUsersHandler(queries))
	server.POST("/signup", SignUpHandler(queries))
	server.POST("/login", LoginHandler(queries))

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.DELETE("/users/:id", DeleteUserHandler(queries))
	authenticated.GET("/todos", GetTodosHandler(queries))
	authenticated.GET("/todos/:id", GetTodoByIDHandler(queries))
	authenticated.POST("/todos", CreateTodoHandler(queries))
	authenticated.DELETE("/todos/:id", DeleteTodoHandler(queries))
}
