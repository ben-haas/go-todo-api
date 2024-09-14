package routes

import (
	"github.com/gin-gonic/gin"
	"todo-api/db"
	"todo-api/middleware"
)

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
