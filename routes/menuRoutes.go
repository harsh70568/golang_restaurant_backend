package routes

import (
	"golang_restaurant_backend/controllers"
	"golang_restaurant_backend/middlewares"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(router *gin.Engine) {
	menu := router.Group("api/v1")
	{
		menu.POST("/createMenu", middlewares.Auntheticate(), controllers.CreateMenu())
		menu.GET("/getMenu/:menuID", middlewares.Auntheticate(), controllers.GetMenu())
		menu.GET("getAllMenus", middlewares.Auntheticate(), controllers.GetAllMenus())
		// menu.PATCH("updateMenu/:menuID", controllers.UpdateMenu())
	}
}
