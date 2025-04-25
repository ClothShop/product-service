package repositories

import (
	"github.com/ClothShop/product-service/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) Create(product *models.Product) error {
	return repo.DB.Create(product).Error
}

func (repo *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := repo.DB.Preload("Images").Find(&products).Error
	return products, err
}

func (repo *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := repo.DB.Preload("Images").First(&product, id).Error
	return &product, err
}

func (repo *ProductRepository) Update(product *models.Product) error {
	return repo.DB.Save(product).Error
}

func (repo *ProductRepository) Delete(id uint) error {
	return repo.DB.Delete(&models.Product{}, id).Error
}
