package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"todo-api/util"
)

var accessSecret = []byte(os.Getenv("JWT_ACCESS_KEY"))

func Authenticate(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	// Check if the header has the correct format: "Bearer <token>"
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer <token>"})
		return
	}

	UserId, Email, err := util.VerifyToken(tokenString, accessSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.Set("userId", UserId)
	c.Set("email", Email)
	c.Next()
}
