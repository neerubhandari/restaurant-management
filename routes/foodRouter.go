package routes

import (
	"github.com/neerubhandari/restaurant-management/controllers"
	"github.com/neerubhandari/restaurant-management/infrastructure"
)

type FoodForRoutes struct {
	handler        infrastructure.Router
	foodController *controllers.FoodController
}

func NewFoodForRoutes(
	handler infrastructure.Router,
	foodController *controllers.FoodController,
) *FoodForRoutes {
	return &FoodForRoutes{
		handler:        handler,
		foodController: foodController,
	}
}

func (fr *FoodForRoutes) Setup() {
	api := fr.handler.Group("/foods")
	api.GET("", fr.foodController.GetFoods)
	api.GET("/:food_id", fr.foodController.GetFood)
	api.POST("", fr.foodController.CreateFood)
	api.PATCH("/:food_id", fr.foodController.UpdateFood)
}
