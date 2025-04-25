package services

import (
	"fmt"
	"github.com/ClothShop/product-service/internal/cache/product_cache"
	"github.com/ClothShop/product-service/internal/dto"
	"github.com/ClothShop/product-service/internal/models"
	"github.com/ClothShop/product-service/internal/repositories"
	"github.com/ClothShop/product-service/internal/repositories/minio"
	"github.com/jinzhu/copier"
	"log"
)

type ProductService struct {
	Repo         *repositories.ProductRepository
	Cache        *product_cache.ProductCache
	ImageService *ImageService
}

func NewProductService(repo *repositories.ProductRepository,
	cache *product_cache.ProductCache,
	imageService *ImageService) *ProductService {
	return &ProductService{Repo: repo, Cache: cache, ImageService: imageService}
}

func (s *ProductService) GetProducts() ([]dto.GetProduct, error) {
	products, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	getProducts := make([]dto.GetProduct, len(products))
	for i, product := range products {
		copier.Copy(&getProducts[i], &product)
	}
	return getProducts, nil
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	cacheKey := s.Cache.BuildCacheKey(id)

	if product, found := s.Cache.GetFromCache(cacheKey); found {
		log.Println("🟢 Cache hit for product_cache:", id)
		return product, nil
	}

	productFromDB, err := s.FindProductById(id)
	if err != nil {
		return nil, err
	}

	s.Cache.SetToCache(cacheKey, productFromDB)
	return productFromDB, nil
}

func (s *ProductService) CreateProduct(productCreate *dto.ProductCreate) error {
	product := &models.Product{}
	if err := copier.Copy(product, productCreate); err != nil {
		return fmt.Errorf("failed to copy product data: %w", err)
	}

	imageUrls := minio.UploadFile(productCreate)
	if imageUrls == nil {
		return fmt.Errorf("failed to upload product images")
	}
	for i := range product.Images {
		product.Images[i] = models.Image{URL: imageUrls[i]}
	}

	if err := s.Repo.Create(product); err != nil {
		return err
	}

	s.Cache.SetToCache(s.Cache.BuildCacheKey(product.ID), product)
	return nil
}

func (s *ProductService) UpdateProduct(productUpdate *dto.ProductUpdate) error {
	existingProduct, err := s.FindProductById(productUpdate.ID)
	if err != nil {
		return err
	}

	if err := copier.Copy(existingProduct, productUpdate); err != nil {
		return fmt.Errorf("failed to copy updated data: %w", err)
	}

	if err := s.Repo.Update(existingProduct); err != nil {
		return err
	}

	s.Cache.SetToCache(s.Cache.BuildCacheKey(existingProduct.ID), existingProduct)
	return nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	if _, err := s.FindProductById(id); err != nil {
		return err
	}

	if err := s.Repo.Delete(id); err != nil {
		return err
	}
	s.Cache.DeleteFromCache(s.Cache.BuildCacheKey(id))
	return nil
}

func (s *ProductService) FindProductById(id uint) (*models.Product, error) {
	return s.Repo.GetByID(id)
}
