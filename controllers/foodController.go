package controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/neerubhandari/restaurant-management/infrastructure"
	"github.com/neerubhandari/restaurant-management/models"
)

type FoodController struct {
	infrastructure.Database
}

func NewFoodController(db infrastructure.Database) *FoodController {
	return &FoodController{
		db,
	}
}

func (fc *FoodController) Migrate() error {
	log.Print("[Food Table Controller]...Migrate")
	err := fc.AutoMigrate(&models.Food{})
	if err != nil {
		log.Print("[Food Table Controller]...Migration Error")
	}
	return err
}

var validate = validator.New()

func (fc *FoodController) GetFoods(c *gin.Context) {
	var returnFood []gin.H
	var food []models.Food

	pageSize, err := strconv.Atoi(c.Query("PageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	offSet := (page - 1) * pageSize
	var count int64
	fc.DB.Model(&models.Food{}).Count(&count)
	if err := infrastructure.NewDatabase().DB.Offset(offSet).Limit(pageSize).Order("created_at desc").Find(&food).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing food items"})
		return
	}
	offSet, err = strconv.Atoi(c.Query("offSet"))
	result := fc.DB.Find(&food)
	if result.Error != nil {
		msg := "Food Not found"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	for _, foodItem := range food {
		returnFood = append(returnFood, gin.H{
			"id":      foodItem.ID,
			"name":    foodItem.Name,
			"price":   foodItem.Price,
			"image":   foodItem.FoodImage,
			"menu_id": foodItem.MenuId,
		})
	}

	c.JSON(http.StatusOK, gin.H{"food": returnFood, "count": count})
}

func (fc *FoodController) GetFood(c *gin.Context) {
	foodIdStr := c.Param("food_id")
	var food models.Food
	foodId, err := strconv.Atoi(foodIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu id"})
		return
	}
	food.ID = uint(foodId)

	if err := fc.DB.First(&food, foodId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error occured while fetching the food item"})
		return
	}

	c.JSON(http.StatusOK, food)

}

func (fc *FoodController) CreateFood(c *gin.Context) {
	var menu models.Menu
	var food models.Food
	if err := c.BindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(food)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	if err := fc.DB.First(&menu, food.MenuId).Error; err != nil {
		msg := fmt.Sprintf("Menu was not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	var num = toFixed(*food.Price, 2)
	food.Price = &num

	if err := fc.DB.Create(&food).Error; err != nil {
		msg := fmt.Sprintf("food Item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, food)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func (fc *FoodController) UpdateFood(c *gin.Context) {
	var food models.Food
	var menu models.Menu

	if err := c.BindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foodIdStr := c.Param("food_id")
	foodId, err := strconv.Atoi(foodIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu id"})
		return
	}
	food.ID = uint(foodId)
	if food.MenuId != nil {
		if err := fc.DB.First(&menu, food.MenuId).Error; err != nil {
			msg := fmt.Sprintf("message: menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	if err := fc.DB.Model(&food).Updates(food).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Food update failed"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Food updated successfully"})

}
