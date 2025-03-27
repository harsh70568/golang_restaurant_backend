package middlewares

import (
	"golang_restaurant_backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

/* Function to check if token is valid and is not expired */
func Auntheticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing access token"})
			c.Abort()
			return
		}

		tk, err := utils.ValidateToken(token)
		if err != nil {
			if err.Error() == "token has expired" {
				refreshToken, err := c.Cookie("refresh_token")
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing refresh token"})
					c.Abort()
					return
				}

				refreshtk, err := utils.ValidateToken(refreshToken)
				if err != nil { /* Refresh token is either invalid or expired */
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
					c.Abort()
					return
				} else {
					/* Generate new tokens */
					claims, ok := refreshtk.Claims.(jwt.MapClaims)
					if !ok {
						c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
						c.Abort()
						return
					}

					email := claims["email"].(string)
					newAccessToken, err := utils.GenerateToken(email)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
						return
					}
					newRefreshToken, err := utils.GenerateRefreshToken(email)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new refresh token"})
						return
					}

					c.SetCookie("token", newAccessToken, int(24*time.Hour.Seconds()), "/", "localhost", false, true)
					c.SetCookie("refresh_token", newRefreshToken, int(24*7*time.Hour.Seconds()), "/", "localhost", false, true)

					c.Next()
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}
		}

		/* access token is valid */
		claims, ok := tk.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		email := claims["email"].(string)
		c.Set("email", email)
		c.Next()
	}
}
