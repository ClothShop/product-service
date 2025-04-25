package models

import (
	"time"
)

type Image struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `json:"product_id"`
	URL       string    `gorm:"type:text;not null" json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
