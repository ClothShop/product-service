package main

import (
	"github.com/ClothShop/product-service/internal/config/minio"
	"log"
	"os"

	"github.com/ClothShop/product-service/internal/cache"
	"github.com/ClothShop/product-service/internal/config/db"
	"github.com/ClothShop/product-service/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.InitDB()
	minio.InitMinio()
	cache.InitRedis()

	r := routes.RegisterRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Server started on: ", port)
	log.Fatal(r.Run(":" + port))
}
