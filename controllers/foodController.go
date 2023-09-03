package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/neerubhandari/restaurant-management/database"
	"github.com/neerubhandari/restaurant-management/models"
)

var db = database.Connection()
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		db.Model(&models.Food{}).Count(&count)
		if err := db.Offset(offSet).Limit(pageSize).Order("created_at desc").Find(&food).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing food items"})
			return
		}
		offSet, err = strconv.Atoi(c.Query("offSet"))
		result := db.Find(&food)
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
}
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodIdStr := c.Param("food_id")
		var food models.Food
		foodId, err := strconv.Atoi(foodIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu id"})
			return
		}
		food.ID = uint(foodId)

		if err := db.First(&food, foodId).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "error occured while fetching the food item"})
			return
		}

		c.JSON(http.StatusOK, food)

	}
}
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		err := db.First(&menu, food.MenuId)
		if err != nil {
			msg := fmt.Sprintf("Menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result := db.Create(&food)
		if result.Error != nil {
			msg := fmt.Sprintf("food Item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
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
			if err := db.First(&menu, food.MenuId).Error; err != nil {
				msg := fmt.Sprintf("message: menu was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
		}
		if err := db.Model(&food).Updates(food).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Food update failed"})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Food updated successfully"})

	}
}
