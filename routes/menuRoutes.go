package routes

import (
	"golang_restaurant_backend/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(router *gin.Engine) {
	menu := router.Group("api/v1")
	{
		menu.POST("/createMenu", controllers.CreateMenu())
		menu.GET("/getMenu/:menuID", controllers.GetMenu())
		menu.GET("getAllMenus", controllers.GetAllMenus())
		// menu.PATCH("updateMenu/:menuID", controllers.UpdateMenu())
	}
}
