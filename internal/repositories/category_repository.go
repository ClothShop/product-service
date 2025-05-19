package repositories

import (
	"github.com/ClothShop/product-service/internal/config/db"
	"github.com/ClothShop/product-service/internal/models"
)

func GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	err := db.DB.First(&category, id).Error
	return category, err
}

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := db.DB.Find(&categories).Error
	return categories, err
}
