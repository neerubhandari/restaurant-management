package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/neerubhandari/restaurant-management/database"

	// "github.com/neerubhandari/restaurant-management/middleware"
	"github.com/neerubhandari/restaurant-management/routes"
)

func main() {
	database.Connect()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	// routes.UserRoutes(router)
	// router.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	// routes.TableRoutes(router)
	// routes.OrderRoutes(router)
	// routes.OrderItemsRoutes(router)
	// routes.InvoiceRoutes(router)
	router.Run(":" + port)
}
