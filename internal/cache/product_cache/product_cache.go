package product_cache

import (
	"encoding/json"
	"fmt"
	"github.com/ClothShop/product-service/internal/cache"
	"github.com/ClothShop/product-service/internal/dto/product"
	"log"
	"os"
	"strconv"
	"time"
)

func BuildCacheKey(id uint) string {
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

func GetFromCache(key string) (*product.GetProduct, bool) {
	cached, err := cache.GetCache(key)
	if err != nil || cached == "" {
		return nil, false
	}

	var product product.GetProduct
	if err := json.Unmarshal([]byte(cached), &product); err != nil {
		log.Println("⚠️ Failed to unmarshal cache:", err)
		return nil, false
	}
	return &product, true
}

func SetToCache(key string, product *product.GetProduct) {
	bytes, err := json.Marshal(product)
	if err != nil {
		log.Println("⚠️ Failed to marshal product_cache to cache:", err)
		return
	}
	cache.SetCache(key, string(bytes), cacheTTL)
}

func DeleteFromCache(key string) {
	cache.DeleteCache(key)
}
