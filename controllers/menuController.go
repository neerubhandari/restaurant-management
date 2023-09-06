package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/neerubhandari/restaurant-management/infrastructure"
	"github.com/neerubhandari/restaurant-management/models"
)

type MenuController struct {
	infrastructure.Database
}

func NewMenuController(db infrastructure.Database) *MenuController {
	return &MenuController{
		db,
	}
}

func (mc *MenuController) Migrate() error {
	log.Print("[Menu Table Controller]...Migrate")
	err := mc.AutoMigrate(&models.Menu{})
	if err != nil {
		log.Print("[Menu Table Controller]...Migration Error")
	}
	return err
}

func (mc *MenuController) GetMenus(c *gin.Context) {
	var menu []models.Menu
	var returnMenu []gin.H
	result := mc.DB.Find(&menu)
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

func (mc *MenuController) GetMenu(c *gin.Context) {
	var menu models.Menu
	menuIdStr := c.Param("menu_id")
	menuId, err := strconv.Atoi(menuIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu id"})
		return
	}
	menu.ID = uint(menuId)

	result := mc.DB.First(&menu, menuId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the menu"})
	}

	c.JSON(http.StatusOK, menu)
}

func (mc *MenuController) CreateMenu(c *gin.Context) {
	log.Print("Creating Menus")

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

	if err := mc.DB.Create(&menu).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menu)

}

func inTimeSpan(startDate, endDate, check time.Time) bool {
	return startDate.After(time.Now()) && endDate.After(startDate)
}

func (mc *MenuController) UpdateMenu(c *gin.Context) {
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
		// if err := mc.DB.First(&menu, menuId).Error; err != nil {
		// 	c.JSON(http.StatusNotFound, gin.H{"error": "Failed to update menu"})
		// 	return
		// }
		if err := mc.DB.Model(&menu).Updates(menu).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "menu update failed"})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Menu updated successfully"})
	}

}
