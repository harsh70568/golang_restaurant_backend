package controllers

import (
	"golang_restaurant_backend/db"
	"golang_restaurant_backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var food models.Food
		var menu models.Menu
		if err := c.ShouldBindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		if err := db.DB.Where("menu_id = ?", food.MenuID).Find(&menu).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Menu ID doesn't exists"})
			return
		}
		food.Created_at = time.Now()
		food.Updated_at = time.Now()

		if err := db.DB.Create(&food).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating food"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Food created suscesfully", "food": food})
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodID := c.Param("foodID")
		if foodID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Food ID doesn't exists"})
			return
		}

		var food models.Food
		if err := db.DB.Where("food_id = ?", foodID).Find(&food).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Food ID doesn't exists"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"food": food})
	}
}

func GetAllFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var allFoods []models.Food
		if err := db.DB.Find(&allFoods).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in getting all foods"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"foods": allFoods})
	}
}