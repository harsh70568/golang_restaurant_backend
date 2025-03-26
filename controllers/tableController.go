package controllers

import (
	"golang_restaurant_backend/db"
	"golang_restaurant_backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table
		if err := c.ShouldBindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}
		table.Created_at = time.Now()
		table.Updated_at = time.Now()

		if err := db.DB.Create(table).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting table"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Table created succesfully", "table": table})
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		tableID := c.Param("tableID")
		if tableID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid TableID"})
			return
		}

		var table models.Table
		if err := db.DB.Where("table_id = ?", tableID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Table ID doesn't exists"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"table": table})
	}
}

func GetAllTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var allTables []models.Table
		if err := db.DB.Find(&allTables).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in getting all tables"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"tables": allTables})
	}
}
