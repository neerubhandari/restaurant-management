package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/neerubhandari/restaurant-management/infrastructure"
	"github.com/neerubhandari/restaurant-management/models"
)

type OrderController struct {
	infrastructure.Database
}

func NewOrderController(db infrastructure.Database) *OrderController {
	return &OrderController{
		db,
	}
}

func (oc *OrderController) Migrate() error {
	log.Print("[Order Table Controller]...Migrate")
	err := oc.AutoMigrate(&models.Order{})
	if err != nil {
		log.Print("[Order Table Controller]...Migration Error")
	}
	return err
}

func (oc *OrderController) GetOrders(c *gin.Context) {
	var order []models.Order
	var returnOrder []gin.H
	result := oc.DB.Find(&order)
	if result.Error != nil {
		msg := "error occured while fetching orders"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	for _, menuItem := range order {
		returnOrder = append(returnOrder, gin.H{
			"id":         menuItem.ID,
			"order_date": menuItem.Order_Date,
			"table_id":   menuItem.Table_id,
		})
	}
	c.JSON(http.StatusOK, gin.H{"menu": returnOrder})
}

func (db *OrderController) GetOrder(c *gin.Context) {

}

func (oc *OrderController) CreateOrder(c *gin.Context) {

	var order models.Order
	var table models.Table
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(order)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	if order.Table_id != nil {
		if err := oc.DB.First(&table, order.Table_id).Error; err != nil {
			msg := fmt.Sprintf("message:Table was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}

	result := oc.DB.Create(&order)
	if result.Error != nil {
		msg := fmt.Sprintf("Order Item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (oc *OrderController) UpdateOrder(c *gin.Context) {
	var order models.Order
	var table models.Table
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderIdStr := c.Param("order_id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order id"})
		return
	}
	order.ID = uint(orderId)
	if order.Table_id != nil {
		if err := oc.DB.First(&table, order.Table_id).Error; err != nil {
			msg := fmt.Sprintf("message: table was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
	if err := oc.DB.Model(&order).Updates(order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Order update failed"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})

}
