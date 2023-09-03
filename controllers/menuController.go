package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/neerubhandari/restaurant-management/models"
)

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu []models.Menu
		var returnMenu []gin.H
		result := db.Find(&menu)
		if result.Error != nil {
			msg := "Menu not found"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		for _, menuItem := range menu {
			returnMenu = append(returnMenu, gin.H{
				"id":         menuItem.ID,
				"name":       menuItem.Name,
				"category":   menuItem.Category,
				"start_date": menuItem.Start_Date,
				"end_date":   menuItem.End_Date,
			})
		}
		c.JSON(http.StatusOK, gin.H{"menu": returnMenu})

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		menuIdStr := c.Param("menu_id")
		menuId, err := strconv.Atoi(menuIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu id"})
			return
		}
		menu.ID = uint(menuId)

		result := db.First(&menu, menuId)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the menu"})
		}

		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(menu)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		result := db.Create(&menu)
		if result.Error != nil {
			msg := "Menu was not created"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func inTimeSpan(startDate, endDate, check time.Time) bool {
	return startDate.After(time.Now()) && endDate.After(startDate)
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		menuIdStr := c.Param("menu_id")
		menuId, err := strconv.Atoi(menuIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu id"})
			return
		}
		menu.ID = uint(menuId)

		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
				msg := "kindly retype the time"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			if err := db.First(&menu, menuId).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Failed to update menu"})
				return
			}
			if err := db.Model(&menu).Updates(menu).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "menu update failed"})
			}
			c.JSON(http.StatusOK, gin.H{"message": "Menu updated successfully"})
		}

	}
}
