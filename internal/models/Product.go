package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"product_name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`

	Images []Image `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"images"`
}
