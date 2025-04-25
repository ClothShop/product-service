package main

import (
	"github.com/ClothShop/product-service/internal/cache/product_cache"
	"github.com/ClothShop/product-service/internal/config/minio"
	"log"
	"os"

	"github.com/ClothShop/product-service/internal/cache"
	"github.com/ClothShop/product-service/internal/config/db"
	"github.com/ClothShop/product-service/internal/handlers"
	"github.com/ClothShop/product-service/internal/repositories"
	"github.com/ClothShop/product-service/internal/routes"
	"github.com/ClothShop/product-service/internal/services"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.InitDB()
	minio.InitMinio()
	cache.InitRedis()

	imageRepository := *repositories.NewImageRepository(db.DB)
	imageService := services.NewImageService(imageRepository)
	productCache := product_cache.NewProductCache()
	repo := repositories.NewProductRepository(db.DB)
	service := services.NewProductService(repo, productCache, imageService)
	handler := handlers.ProductHandler{Service: service}

	r := routes.RegisterRoutes(&handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Server started on: ", port)
	log.Fatal(r.Run(":" + port))
}
