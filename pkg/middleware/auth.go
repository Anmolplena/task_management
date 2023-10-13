package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	adminRole   = "admin"
	defaultRole = "default"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Parse the token without a secret key (insecure for demonstration purposes)
		token, _, err := new(jwt.Parser).ParseUnverified(tokenParts[1], jwt.MapClaims{})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract the role from the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in token"})
			c.Abort()
			return
		}

		// Set the role in the Gin context for later use
		c.Set("userRole", role)
		// fmt.Println(role)

		// Continue with the request
		c.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not determine user role"})
			c.Abort()
			return
		}

		// Access control based on user role
		if userRole != adminRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access only"})
			c.Abort()
			return
		}

		// Continue with the request
		c.Next()
	}
}

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not determine user role"})
			c.Abort()
			return
		}

		// Access control based on user role
		if userRole != defaultRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "User access only"})
			c.Abort()
			return
		}

		// Continue with the request
		c.Next()
	}
}
