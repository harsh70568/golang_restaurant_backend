package routes

import (
	"golang_restaurant_backend/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(router *gin.Engine) {
	Table := router.Group("api/v1")
	{
		Table.POST("/createTable", controllers.CreateTable())
		Table.GET("/getTable/:tableID", controllers.GetTable())
		Table.GET("/getAllTables", controllers.GetAllTables())
		// Table.PATCH("/updateTables", controllers.UpdateTables())
	}
}
