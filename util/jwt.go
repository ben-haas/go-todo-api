package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var accessSecret = []byte(os.Getenv("JWT_ACCESS_KEY"))
var refreshSecret = []byte(os.Getenv("JWT_REFRESH_KEY"))

const accessTokenExpiration = time.Hour
const refreshTokenExpiration = 7 * 24 * time.Hour

type Claims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateAccessToken generates a new JWT token for the authenticated user
func GenerateAccessToken(userID int64, email string) (string, error) {
	expirationTime := time.Now().Add(accessTokenExpiration)
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(accessSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateRefreshToken generates a new JWT token for the authenticated user
func GenerateRefreshToken(userID int64, email string) (string, error) {
	expirationTime := time.Now().Add(refreshTokenExpiration)
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(refreshSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies the JWT token and returns the user ID and email if valid
func VerifyToken(tokenString string, secret []byte) (int64, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, claims.Email, nil
	}

	return 0, "", errors.New("invalid token")
}

// RefreshAccessToken generates a new access token using the refresh token
func RefreshAccessToken(refreshToken string) (string, error) {
	UserId, Email, err := VerifyToken(refreshToken, refreshSecret)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("invalid token")
	}

	newAccessToken, err := GenerateAccessToken(UserId, Email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return newAccessToken, nil
}
