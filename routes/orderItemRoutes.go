package routes

import (
	"golang_restaurant_backend/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItem(router *gin.Engine) {
	orderItem := router.Group("api/v1")
	{
		orderItem.POST("/orderItem", controllers.OrderItem())
		orderItem.GET("/getOrderItem/:orderItemID", controllers.GetOrderItem())
		orderItem.GET("/getAllOrderItem", controllers.GetAllOrderItems())
		// orderItem.GET("/getOrderItemOrder/:orderID", controllers.getOrderItemOrder())
	}
}
