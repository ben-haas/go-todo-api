package routes

import (
	"github.com/gin-gonic/gin"
	"todo-api/internal/db"
	"todo-api/middleware"
)

// TODO add admin routes
// TODO handle token refresh client side

// RegisterRoutes registers all the application's routes with Gin.
func RegisterRoutes(server *gin.Engine, queries *db.Queries) {
	v1 := server.Group("/v1")

	v1.GET("/users", GetUsersHandler(queries))
	v1.POST("/signup", SignUpHandler(queries))
	v1.POST("/login", LoginHandler(queries))
	v1.POST("/refresh", RefreshTokenHandler())

	authenticated := v1.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.DELETE("/users/:id", DeleteUserHandler(queries))
	authenticated.GET("/todos", GetTodosHandler(queries))
	authenticated.GET("/todos/:id", GetTodoByIDHandler(queries))
	authenticated.POST("/todos", CreateTodoHandler(queries))
	authenticated.DELETE("/todos/:id", DeleteTodoHandler(queries))
}
