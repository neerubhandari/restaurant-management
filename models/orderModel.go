package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Order_Date time.Time `json:"order_date" validate:"required"`
	Table_id   *string   `json:"table_id"  validate:"required"`
}
