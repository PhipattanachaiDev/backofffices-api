package middlewares

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	response "ezview.asia/ezview-web/ezview-lite-back-office/types/responses"
	"github.com/gin-gonic/gin"
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

// Claims represents the JWT claims structure
type Claims struct {
	UserId   int             `json:"user_id"`
	Username string          `json:"username"`
	RoleId   int             `json:"role_id"`
	RoleName string          `json:"role_name"`
	Access   json.RawMessage `json:"access"`
	jwt.RegisteredClaims
}

// JWTMiddleware is a middleware function to handle JWT authentication
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "Authorization header required")
			c.Abort()
			return
		}

		// Extract token from the header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &Claims{}

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		switch {
		case token.Valid:
			c.Set("claims", claims)
			c.Set("user_id", claims.UserId)
			c.Next()
		case errors.Is(err, jwt.ErrTokenMalformed):
			response.Unauthorized(c, "Token is invalid")
			c.Abort()
			return
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			response.Unauthorized(c, "Token is invalid")
			c.Abort()
			return
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			response.Unauthorized(c, "Token has expired")
			c.Abort()
			return
		default:
			response.Unauthorized(c, "Token is invalid")
			c.Abort()
			return
		}
	}
}

// APIVersionMiddleware is a middleware function for API versioning
func APIVersionMiddleware(c *gin.Context) {
	apiVersion := c.GetHeader("API-Version")
	if apiVersion != "" {
		c.Set("API-Version", apiVersion)
	} else {
		c.Set("API-Version", "v1") // Default version if not specified
	}
	c.Next()
}
