package middleware

import (
	"fmt"
	"net/http"

	"github.com/25Kamalesh/go_todo_api/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Debug: Print all cookies
		fmt.Println("=== Auth Middleware Debug ===")
		fmt.Println("All cookies:", c.Request.Cookies())

		tokenString, err := c.Cookie("token")
		if err != nil {
			fmt.Println("Error getting cookie:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No token cookie found"})
			return
		}

		fmt.Println("Token found:", tokenString[:20]+"...") // Print first 20 chars

		// 2. Validate token
		userID, err := auth.ValidateToken(tokenString, secret)
		if err != nil {
			fmt.Println("Token validation failed:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		fmt.Println("Token valid, user_id:", userID)

		// 3. Store user_id in context
		c.Set("user_id", userID)

		// 4. Continue to next handler
		c.Next()
	}
}
