package main

import (
	"github.com/neerubhandari/restaurant-management/bootstrap"
	"github.com/neerubhandari/restaurant-management/utils"
)

func main() {

	// database.Connect()
	// if err := database.DB.AutoMigrate(&models.Menu{}); err != nil {
	// 	log.Fatalf("Error during migration: %v", err)
	// }
	utils.LoadEnv()
	bootstrap.WebApp()
	// fmt.Println("Migration successful!")

	// router := gin.Default()
	// router.Use(gin.Logger())
	// routes.UserRoutes(router)
	// router.Use(middleware.Authentication())
	// routes.FoodRoutes(router)
	// routes.MenuRoutes(router)
	// routes.TableRoutes(router)
	// routes.OrderRoutes(router)
	// routes.OrderItemsRoutes(router)
	// routes.InvoiceRoutes(router)

}
