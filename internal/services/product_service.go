package services

import (
	"fmt"
	"github.com/ClothShop/product-service/internal/cache/product_cache"
	"github.com/ClothShop/product-service/internal/dto/product"
	"github.com/ClothShop/product-service/internal/models"
	"github.com/ClothShop/product-service/internal/repositories"
	"github.com/jinzhu/copier"
	"log"
	"os"
)

func GetProducts() ([]product.GetProduct, error) {
	products, err := repositories.GetAll()
	if err != nil {
		return nil, err
	}
	productDtos := make([]product.GetProduct, len(products))
	for i, product := range products {
		copier.Copy(&productDtos[i], &product)
		urls := make([]string, len(product.Images))
		for j, img := range product.Images {
			urls[j] = img.URL
		}
		productDtos[i].Images = urls
	}

	return productDtos, nil
}

func GetProduct(id uint) (*product.GetProduct, error) {
	cacheKey := product_cache.BuildCacheKey(id)
	if product, found := product_cache.GetFromCache(cacheKey); found {
		log.Println("🟢 Cache hit for product_cache:", id)
		return product, nil
	}

	productFromDB, err := FindProductById(id)
	if err != nil {
		return nil, err
	}
	log.Println("productFromDB:", productFromDB)
	var productDto product.GetProduct
	copier.Copy(&productDto, &productFromDB)
	urls := make([]string, len(productFromDB.Images))
	for j, img := range productFromDB.Images {
		urls[j] = img.URL
	}
	productDto.Images = urls

	go product_cache.SetToCache(cacheKey, &productDto)
	return &productDto, nil
}

func CreateProduct(productCreate *product.Create) (*product.GetProduct, error) {
	category, err := GetCategoryById(productCreate.CategoryID)
	if err != nil {
		return nil, err
	}
	productEntity := &models.Product{}
	if err := copier.Copy(productEntity, productCreate); err != nil {
		return nil, fmt.Errorf("failed to copy product data: %w", err)
	}

	productEntity.CategoryID = category.ID
	imageUrls := repositories.UploadFile(productCreate)
	if imageUrls == nil {
		return nil, fmt.Errorf("failed to upload product images")
	}
	urls := make([]string, len(imageUrls))
	for i := range productEntity.Images {
		productEntity.Images[i] = models.Image{URL: "http://" + os.Getenv("MINIO_ENDPOINT") + "/" + os.Getenv("MINIO_BUCKET") + "/" + imageUrls[i]}
		urls[i] = productEntity.Images[i].URL
	}

	if err := repositories.Create(productEntity); err != nil {
		return nil, err
	}

	var productDto product.GetProduct
	copier.Copy(&productDto, &productEntity)
	productDto.Images = urls
	go product_cache.SetToCache(product_cache.BuildCacheKey(productEntity.ID), &productDto)
	return &productDto, nil
}

func UpdateProduct(id uint, productUpdate *product.Update) error {
	existingProduct, err := FindProductById(id)
	if err != nil {
		return err
	}

	if productUpdate.CategoryID != nil {
		existingCategory, err := GetCategoryById(*productUpdate.CategoryID)
		if err != nil {
			return err
		}
		existingProduct.Category = existingCategory
	}

	if err := copier.Copy(existingProduct, productUpdate); err != nil {
		return fmt.Errorf("failed to copy updated data: %w", err)
	}
	existingProduct.ID = id
	if err := repositories.Update(existingProduct); err != nil {
		return err
	}

	var productDto product.GetProduct
	copier.Copy(&productDto, &existingProduct)
	urls := make([]string, len(existingProduct.Images))
	for j, img := range existingProduct.Images {
		urls[j] = img.URL
	}
	productDto.Images = urls
	go product_cache.SetToCache(product_cache.BuildCacheKey(existingProduct.ID), &productDto)
	copier.Copy(&productUpdate, &productDto)
	return nil
}

func DeleteProduct(id uint) error {
	if _, err := FindProductById(id); err != nil {
		return err
	}

	if err := repositories.Delete(id); err != nil {
		return err
	}
	go product_cache.DeleteFromCache(product_cache.BuildCacheKey(id))
	return nil
}

func FindProductById(id uint) (*models.Product, error) {
	return repositories.GetByID(id)
}
