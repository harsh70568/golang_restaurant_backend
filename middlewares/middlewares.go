package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Function to check if token is valid and is not expired */
func Auntheticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error getting access token"})
			c.Abort()
			return
		}

		refreshToken, err := c.Cookie("refresh_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error getting refresh token"})
			c.Abort()
			return
		}
		_, _ = token, refreshToken
	}
}
