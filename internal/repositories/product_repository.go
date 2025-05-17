package repositories

import (
	"github.com/ClothShop/product-service/internal/config/db"
	"github.com/ClothShop/product-service/internal/models"
)

func Create(product *models.Product) error {
	return db.DB.Create(product).Error
}

func GetAll() ([]models.Product, error) {
	var products []models.Product
	err := db.DB.Preload("Images").Find(&products).Error
	return products, err
}

func GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := db.DB.Preload("Images").First(&product, id).Error
	return &product, err
}

func Update(product *models.Product) error {
	return db.DB.Save(product).Error
}

func Delete(id uint) error {
	return db.DB.Delete(&models.Product{}, id).Error
}
