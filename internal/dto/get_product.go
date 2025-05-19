package dto

import (
	"time"
)

type GetProduct struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Price          float64   `json:"price"`
	CompareAtPrice float64   `json:"compare_at_price"`
	CreatedAt      time.Time `json:"created_at"`
	Images         []string  `json:"images"`
	Category       string    `json:"category"`
}
