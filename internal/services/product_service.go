package services

import (
	"fmt"
	"github.com/ClothShop/product-service/internal/cache/product_cache"
	"github.com/ClothShop/product-service/internal/dto"
	"github.com/ClothShop/product-service/internal/models"
	"github.com/ClothShop/product-service/internal/repositories"
	"github.com/jinzhu/copier"
	"log"
	"os"
)

func GetProducts() ([]dto.GetProduct, error) {
	products, err := repositories.GetAll()
	if err != nil {
		return nil, err
	}
	productDtos := make([]dto.GetProduct, len(products))
	for i, product := range products {
		copier.Copy(&productDtos[i], &product)
		productDtos[i].Category = product.Category.Name
		urls := make([]string, len(product.Images))
		for j, img := range product.Images {
			urls[j] = img.URL
		}
		productDtos[i].Images = urls
	}

	return productDtos, nil
}

func GetProduct(id uint) (*dto.GetProduct, error) {
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
	var productDto dto.GetProduct
	copier.Copy(&productDto, &productFromDB)
	productDto.Category = productFromDB.Category.Name
	urls := make([]string, len(productFromDB.Images))
	for j, img := range productFromDB.Images {
		urls[j] = img.URL
	}
	productDto.Images = urls

	go product_cache.SetToCache(cacheKey, &productDto)
	return &productDto, nil
}

func CreateProduct(productCreate *dto.ProductCreate) (*dto.GetProduct, error) {
	category, err := GetCategoryById(productCreate.CategoryID)
	if err != nil {
		return nil, err
	}
	product := &models.Product{}
	if err := copier.Copy(product, productCreate); err != nil {
		return nil, fmt.Errorf("failed to copy product data: %w", err)
	}

	product.CategoryID = category.ID
	imageUrls := repositories.UploadFile(productCreate)
	if imageUrls == nil {
		return nil, fmt.Errorf("failed to upload product images")
	}
	urls := make([]string, len(imageUrls))
	for i := range product.Images {
		product.Images[i] = models.Image{URL: "http://" + os.Getenv("MINIO_ENDPOINT") + "/" + os.Getenv("MINIO_BUCKET") + "/" + imageUrls[i]}
		urls[i] = product.Images[i].URL
	}

	if err := repositories.Create(product); err != nil {
		return nil, err
	}

	var productDto dto.GetProduct
	copier.Copy(&productDto, &product)
	productDto.Category = category.Name
	productDto.Images = urls
	go product_cache.SetToCache(product_cache.BuildCacheKey(product.ID), &productDto)
	return &productDto, nil
}

func UpdateProduct(productUpdate *dto.ProductUpdate) error {
	existingProduct, err := FindProductById(productUpdate.ID)
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

	if err := repositories.Update(existingProduct); err != nil {
		return err
	}

	var productDto dto.GetProduct
	copier.Copy(&productDto, &existingProduct)
	productDto.Category = existingProduct.Category.Name
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
