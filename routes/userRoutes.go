package routes

import (
	"golang_restaurant_backend/controllers"
	"golang_restaurant_backend/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("api/v1/users")
	{
		user.POST("/login", controllers.Login())
		user.POST("/signup", controllers.SignUp())
		user.GET("/getUser/:userID", middlewares.Auntheticate(), controllers.GetUser())
		user.GET("getAllUsers", middlewares.Auntheticate(), controllers.GetAllUsers())
	}
}
