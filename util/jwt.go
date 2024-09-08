package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token for the authenticated user
func GenerateJWT(userID int64, email string) (string, error) {

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // Token valid for 72 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies the JWT token and returns the user ID and email if valid
func VerifyToken(tokenString string) (int64, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, claims.Email, nil
	}

	return 0, "", errors.New("invalid token")
}
