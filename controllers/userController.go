package controllers

import (
	"golang_restaurant_backend/db"
	"golang_restaurant_backend/models"
	"golang_restaurant_backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request", "detail": err.Error()})
			return
		}

		/* Check if email does exists or not */
		var existingUser models.User
		if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}

		/* Hash the password */
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error hashing password"})
			return
		}

		user.Password = hashedPassword
		if err := db.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Inserting data in database"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"success": "true", "id": user.ID})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		if user.Email == "" || user.Password == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		/* Check if entered email does exists or not */
		var existingUser models.User
		if err := db.DB.Where("email = ?", user.Email).Find(&existingUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email does not exists"})
			return
		}

		/* Check if the password entered matched or not */
		if err := utils.ComparePassword(existingUser.Password, user.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Password is incorrect"})
			return
		}

		token, err := utils.GenerateToken(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		refresh_token, err := utils.GenerateRefreshToken(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating refresh token"})
			return
		}

		c.SetCookie("token", token, int(24*time.Hour.Seconds()), "/", "localhost", false, true)
		c.SetCookie("refresh_token", refresh_token, int(24*7*time.Hour.Seconds()), "/", "localhost", false, true)

		c.JSON(http.StatusCreated, gin.H{
			"user":          existingUser,
			"token":         token,
			"refresh_token": refresh_token,
		})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userID")
		if userID == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "userId is null"})
			return
		}

		var existingUser models.User
		if err := db.DB.Where("id = ?", userID).Find(&existingUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": existingUser})
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var allUsers []models.User

		if err := db.DB.Find(&allUsers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in getting all users"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"users": allUsers})
	}
}
