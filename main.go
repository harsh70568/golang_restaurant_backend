package main

import (
	"golang_restaurant_backend/db"
	"golang_restaurant_backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	/* Connect to database */
	db.ConnectDB()

	/* Migrate the databases */
	db.MigrateDatabase()

	/* Setting up router */
	router := gin.Default()

	/* Setting up routes */
	routes.UserRoutes(router)
	routes.MenuRoutes(router)
	routes.FoodRoutes(router)
	routes.TableRoutes(router)
	routes.Order(router)
	routes.OrderItem(router)

	/* Running the server */
	err := router.Run(db.ServerPort)
	if err != nil {
		log.Fatalf("Error in starting the server")
	}

}
