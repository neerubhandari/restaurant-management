package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/neerubhandari/restaurant-management/database"
	"github.com/neerubhandari/restaurant-management/models"

	// "github.com/neerubhandari/restaurant-management/middleware"
	"github.com/neerubhandari/restaurant-management/routes"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	database.Connect()
	if err := database.DB.AutoMigrate(&models.Menu{}); err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	fmt.Println("Migration successful!")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.Default()
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
