package product_cache

import (
	"encoding/json"
	"fmt"
	"github.com/ClothShop/product-service/internal/cache"
	"github.com/ClothShop/product-service/internal/models"
	"log"
	"os"
	"strconv"
	"time"
)

type ProductCache struct {
}

func NewProductCache() *ProductCache {
	return &ProductCache{}
}

func (s *ProductCache) BuildCacheKey(id uint) string {
	var ttlStr = os.Getenv("PRODUCT_CACHE_TTL")
	var ttlInt, err = strconv.Atoi(ttlStr)
	if err != nil {
		log.Printf("⚠️ Invalid PRODUCT_CACHE_TTL: %v, using default (5 min)", err)
		ttlInt = 5
	}
	cacheTTL = time.Duration(ttlInt) * time.Minute
	return fmt.Sprintf("product_cache:%d", id)
}

var cacheTTL time.Duration

func (s *ProductCache) GetFromCache(key string) (*models.Product, bool) {
	cached, err := cache.GetCache(key)
	if err != nil || cached == "" {
		return nil, false
	}

	var product models.Product
	if err := json.Unmarshal([]byte(cached), &product); err != nil {
		log.Println("⚠️ Failed to unmarshal cache:", err)
		return nil, false
	}
	return &product, true
}

func (s *ProductCache) SetToCache(key string, product *models.Product) {
	bytes, err := json.Marshal(product)
	if err != nil {
		log.Println("⚠️ Failed to marshal product_cache to cache:", err)
		return
	}
	cache.SetCache(key, string(bytes), cacheTTL)
}

func (s *ProductCache) DeleteFromCache(key string) {
	cache.DeleteCache(key)
}
