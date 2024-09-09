package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement your authentication logic here
		// For example, check for a valid JWT token

		// If authentication fails:
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		// c.Abort()
		// return

		c.Next()
	}
}
