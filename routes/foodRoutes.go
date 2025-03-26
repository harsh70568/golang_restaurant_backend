package routes

import (
	"golang_restaurant_backend/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(router *gin.Engine) {
	food := router.Group("api/v1/food")
	{
		food.POST("/createFood", controllers.CreateFood())
		food.GET("/getFood/:foodID", controllers.GetFood())
		food.GET("/getAllFoods", controllers.GetAllFoods())
	}
}

