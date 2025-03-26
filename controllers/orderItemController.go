package controllers

import (
	"golang_restaurant_backend/db"
	"golang_restaurant_backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderItemPack struct {
	TableID    uint
	OrderItems []models.OrderItem
}

func OrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderItemPack OrderItemPack
		if err := c.ShouldBindJSON(&orderItemPack); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		var order models.Order
		order.Order_date = time.Now()
		order.TableID = orderItemPack.TableID

		/* Create an order in the database to get orderID */
		if err := db.DB.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting order"})
			return
		}

		var orderItemInserted []models.OrderItem
		for _, orderItem := range orderItemPack.OrderItems {
			orderItem.OrderID = order.ID
			orderItem.Created_at = time.Now()
			orderItem.Updated_at = time.Now()
			// orderItem.UnitPrice = orderItem.UnitPrice

			orderItemInserted = append(orderItemInserted, orderItem)
		}

		if err := db.DB.Create(&orderItemInserted).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert order items"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"order_items": orderItemInserted})
	}
}

func GetAllOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderItems []models.OrderItem
		if err := db.DB.Find(&orderItems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting all orderItems"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"all_order_items": orderItems})
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemID := c.Param("orderItemID")
		if orderItemID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "orderItemID is not valid"})
			return
		}

		var orderItem models.OrderItem
		if err := db.DB.Where("order_item_id = ?", orderItemID).Find(&orderItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find orderItem"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"order_item": orderItem})
	}
}
