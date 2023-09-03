package models

import (
	"time"
)

type Note struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Text       string    `json:"text"`
	Title      string    `json:"title"`
	Created_at time.Time `json:"created_at" `
	Updated_at time.Time `json:"updated_at" `
	Note_id    string    `json:"note_id" `
}
