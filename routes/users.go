package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-api/internal/db"
	"todo-api/util"
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

// SignUpHandler registers a new user
func SignUpHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user db.CreateUserParams

		// Bind JSON body to the user struct
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		// Hash the password
		hashedPassword, err := util.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		params := db.CreateUserParams{
			Email:    user.Email,
			Password: hashedPassword,
		}

		err = queries.CreateUser(context.Background(), params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		// Send success response
		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}

// DeleteUserHandler deletes a user in the database
func DeleteUserHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		err = queries.DeleteUser(context.Background(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{"message": "User deleted successfully"})
	}
}

// LoginHandler authenticates a user and returns a JWT token
func LoginHandler(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user db.CreateUserParams

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		dbUser, err := queries.GetUserByEmail(context.Background(), user.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Compare the provided password with the hashed password in the database
		err = util.CheckPasswordHash(user.Password, dbUser.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		accessToken, err := util.GenerateAccessToken(dbUser.ID, dbUser.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		refreshToken, err := util.GenerateRefreshToken(dbUser.ID, dbUser.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
	}
}

// RefreshTokenHandler generates a new access token using a refresh token
func RefreshTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken := c.Request.Header.Get("Refresh-Token")
		if refreshToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is missing"})
			return
		}

		newAccessToken, err := util.RefreshAccessToken(refreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
	}
}
