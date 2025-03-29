package routes

import (
	"golang_restaurant_backend/controllers"
	"golang_restaurant_backend/middlewares"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(router *gin.Engine) {
	food := router.Group("api/v1/food")
	{
		food.POST("/createFood", middlewares.Auntheticate(), controllers.CreateFood())
		food.GET("/getFood/:foodID", middlewares.Auntheticate(), controllers.GetFood())
		food.GET("/getAllFoods", middlewares.Auntheticate(), controllers.GetAllFoods())
	}
}
