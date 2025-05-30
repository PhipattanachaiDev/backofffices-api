package jwt

import (
	"fmt"
	"log"
	"os"

	"ezview.asia/ezview-web/ezview-lite-back-office/middlewares"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Load environment variables from .env file
func init() {
	if os.Getenv("GIN_MODE") != "release" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
}

// Retrieve JWT key from environment variable
var jwtKey = []byte(os.Getenv("JWT_KEY"))

// VerifyRefreshToken verifies and parses the refresh token
func VerifyRefreshToken(tokenString string) (*middlewares.Claims, error) {
	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &middlewares.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*middlewares.Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
