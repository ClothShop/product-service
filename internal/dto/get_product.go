package dto

import (
	"time"
)

type GetProduct struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	Images      []string  `json:"images"`
}
