package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name       string     `json:"name" validate:"required"`
	Category   string     `json:"category" validate:"required"`
	Start_Date *time.Time `json:"start_date" `
	End_Date   *time.Time `json:"end_date"`
}
