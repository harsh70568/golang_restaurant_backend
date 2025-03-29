package routes

import (
	"golang_restaurant_backend/controllers"
	"golang_restaurant_backend/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderItem(router *gin.Engine) {
	orderItem := router.Group("api/v1")
	{
		orderItem.POST("/orderItem", middlewares.Auntheticate(), controllers.OrderItem())
		orderItem.GET("/getOrderItem/:orderItemID", middlewares.Auntheticate(), controllers.GetOrderItem())
		orderItem.GET("/getAllOrderItem", middlewares.Auntheticate(), controllers.GetAllOrderItems())
		// orderItem.GET("/getOrderItemOrder/:orderID", controllers.getOrderItemOrder())
	}
}
