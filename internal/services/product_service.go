package services

import (
	"fmt"
	"github.com/ClothShop/product-service/internal/cache/product_cache"
	"github.com/ClothShop/product-service/internal/dto"
	"github.com/ClothShop/product-service/internal/models"
	"github.com/ClothShop/product-service/internal/repositories"
	"github.com/jinzhu/copier"
	"log"
)

func GetProducts() ([]dto.GetProduct, error) {
	products, err := repositories.GetAll()
	if err != nil {
		return nil, err
	}
	getProducts := make([]dto.GetProduct, len(products))
	for i, product := range products {
		copier.Copy(&getProducts[i], &product)
	}
	return getProducts, nil
}

func GetProduct(id uint) (*models.Product, error) {
	cacheKey := product_cache.BuildCacheKey(id)

	if product, found := product_cache.GetFromCache(cacheKey); found {
		log.Println("🟢 Cache hit for product_cache:", id)
		return product, nil
	}

	productFromDB, err := FindProductById(id)
	if err != nil {
		return nil, err
	}

	product_cache.SetToCache(cacheKey, productFromDB)
	return productFromDB, nil
}

func CreateProduct(productCreate *dto.ProductCreate) error {
	product := &models.Product{}
	if err := copier.Copy(product, productCreate); err != nil {
		return fmt.Errorf("failed to copy product data: %w", err)
	}

	imageUrls := repositories.UploadFile(productCreate)
	if imageUrls == nil {
		return fmt.Errorf("failed to upload product images")
	}
	for i := range product.Images {
		product.Images[i] = models.Image{URL: imageUrls[i]}
	}

	if err := repositories.Create(product); err != nil {
		return err
	}

	product_cache.SetToCache(product_cache.BuildCacheKey(product.ID), product)
	return nil
}

func UpdateProduct(productUpdate *dto.ProductUpdate) error {
	existingProduct, err := FindProductById(productUpdate.ID)
	if err != nil {
		return err
	}

	if err := copier.Copy(existingProduct, productUpdate); err != nil {
		return fmt.Errorf("failed to copy updated data: %w", err)
	}

	if err := repositories.Update(existingProduct); err != nil {
		return err
	}

	product_cache.SetToCache(product_cache.BuildCacheKey(existingProduct.ID), existingProduct)
	return nil
}

func DeleteProduct(id uint) error {
	if _, err := FindProductById(id); err != nil {
		return err
	}

	if err := repositories.Delete(id); err != nil {
		return err
	}
	product_cache.DeleteFromCache(product_cache.BuildCacheKey(id))
	return nil
}

func FindProductById(id uint) (*models.Product, error) {
	return repositories.GetByID(id)
}
