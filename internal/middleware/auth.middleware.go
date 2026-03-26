package middleware

import (
	"net/http"

	"github.com/25Kamalesh/go_todo_api/internal/auth"
	"github.com/gin-gonic/gin"
)


func AuthMiddleware(secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString , err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// 2. Validate token
		userID, err := auth.ValidateToken(tokenString, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 3. Store user_id in context
		c.Set("user_id", userID)

		// 4. Continue to next handler
		c.Next()
	}
}