package controllers

import (
	"golang_restaurant_backend/db"
	"golang_restaurant_backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Order() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		var table models.Table
		if err := db.DB.Where("table_id = ?", order.TableID).Find(&table).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		table.Created_at = time.Now()
		table.Updated_at = time.Now()

		c.JSON(http.StatusOK, gin.H{"order": order})
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("order_id")
		if orderID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error in getting order ID"})
			return
		}

		var order models.Order
		if err := db.DB.Where("order_id = ?", orderID).Find(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Order Id doesn't exists"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Order": order})
	}
}

func GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var allOrders []models.Order
		if err := db.DB.Find(&allOrders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in getting all orders"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"All orders": allOrders})
	}
}
