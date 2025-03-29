package routes

import (
	"golang_restaurant_backend/controllers"
	"golang_restaurant_backend/middlewares"

	"github.com/gin-gonic/gin"
)

func TableRoutes(router *gin.Engine) {
	Table := router.Group("api/v1")
	{
		Table.POST("/createTable", middlewares.Auntheticate(), controllers.CreateTable())
		Table.GET("/getTable/:tableID", middlewares.Auntheticate(), controllers.GetTable())
		Table.GET("/getAllTables", middlewares.Auntheticate(), controllers.GetAllTables())
		// Table.PATCH("/updateTables", controllers.UpdateTables())
	}
}
