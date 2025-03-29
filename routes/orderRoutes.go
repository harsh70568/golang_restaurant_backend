package routes

import (
	"golang_restaurant_backend/controllers"
	"golang_restaurant_backend/middlewares"

	"github.com/gin-gonic/gin"
)

func Order(router *gin.Engine) {
	Order := router.Group("api/v1")
	{
		Order.POST("/order", middlewares.Auntheticate(), controllers.Order())
		Order.GET("/getOrder/:order_id", middlewares.Auntheticate(), controllers.GetOrder())
		Order.GET("/getAllOrders", middlewares.Auntheticate(), controllers.GetAllOrders())
	}
}
