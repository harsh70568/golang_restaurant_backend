package controllers

import (
	"golang_restaurant_backend/db"
	"golang_restaurant_backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		if err := c.ShouldBindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		menu.Created_at = time.Now()
		menu.Updated_at = time.Now()

		if err := db.DB.Create(&menu).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating menu item"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Menu created sucesfully", "id": menu.ID})
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuID := c.Param("menuID")
		if menuID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "MenuID is invalid"})
			return
		}

		var existingMenu models.Menu
		if err := db.DB.Where("menu_id = ?", menuID).Find(&existingMenu).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No menu with requested menu id"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"menu": existingMenu})
	}
}

func GetAllMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var allMenus []models.Menu

		if err := db.DB.Find(&allMenus).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in getting all menus"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"menus": allMenus})
	}
}
