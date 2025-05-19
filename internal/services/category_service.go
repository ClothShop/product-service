package services

import (
	"fmt"
	"github.com/ClothShop/product-service/internal/models"
	"github.com/ClothShop/product-service/internal/repositories"
	"log"
)

func GetCategoryById(id uint) (models.Category, error) {
	category, err := repositories.GetCategoryByID(id)
	if err != nil {
		log.Println("error getting category by id:", err)
		return models.Category{}, fmt.Errorf("category not found: %w", err)
	}
	return category, nil
}
