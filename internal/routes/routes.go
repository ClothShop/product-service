package routes

import (
	"fmt"
	"github.com/ClothShop/product-service/internal/handlers"
	"github.com/ClothShop/product-service/internal/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func RegisterRoutes(handler *handlers.ProductHandler) *gin.Engine {
	apiVersion := os.Getenv("API_VERSION")
	baseURL := fmt.Sprintf("/api/%s/products", apiVersion)

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8888"},
		AllowWildcard:    false,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	product := r.Group(baseURL)
	{
		product.Use(middlewares.AuthMiddleware())
		product.GET("", handler.GetProducts)
		product.GET("/:id", handler.GetProduct)
		product.DELETE("/:id", middlewares.AdminMiddleware(), handler.DeleteProduct)

		product.POST("", middlewares.ProductCreateMiddleware(), middlewares.AdminMiddleware(), handler.CreateProduct)
		product.PUT("", middlewares.ProductUpdateMiddleware(), middlewares.AdminMiddleware(), handler.UpdateProduct)
	}

	return r
}
