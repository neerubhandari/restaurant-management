package models

import (
	"time"
)

type OrderItem struct {
	ID            uint      `gorm:"primarykey" json:"id"`
	Quantity      *string   `json:"order_date" validate:"required,eq=S |eq=M|eq=L"`
	Unit_price    *float64  `json:"unit_price"  validate:"required"`
	Created_at    time.Time `json:"created_at" `
	Updated_at    time.Time `json:"updated_at" `
	Order_id      string    `json:"order_id" validate:"required"`
	Order_item_id string    `json:"order_item_id" `
	Food_id       *string   `json:"food_id"  validate:"required"`
}
