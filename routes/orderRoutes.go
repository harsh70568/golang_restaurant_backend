package routes

import (
	"golang_restaurant_backend/controllers"

	"github.com/gin-gonic/gin"
)

func Order(router *gin.Engine) {
	Order := router.Group("api/v1")
	{
		Order.POST("/order", controllers.Order())
		Order.GET("/getOrder/:order_id", controllers.GetOrder())
		Order.GET("/getAllOrders", controllers.GetAllOrders())
	}
}
