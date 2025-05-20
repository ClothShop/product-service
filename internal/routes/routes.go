package routes

import (
	"fmt"
	"github.com/ClothShop/product-service/internal/dto/product"
	"github.com/ClothShop/product-service/internal/handlers"
	"github.com/ClothShop/product-service/internal/middlewares"
	"github.com/gin-gonic/gin"
	"os"
)

func RegisterRoutes() *gin.Engine {
	apiVersion := os.Getenv("API_VERSION")
	baseURL := fmt.Sprintf("/api/%s", apiVersion)

	r := gin.Default()

	productRoutes := r.Group(baseURL + "/products")
	{
		productRoutes.GET("/", handlers.GetProducts)
		productRoutes.GET("/:id", handlers.GetProduct)
		productRoutes.Use(middlewares.AuthMiddleware())
		productRoutes.DELETE("/:id", middlewares.AdminMiddleware(), handlers.DeleteProduct)

		productRoutes.POST("/", middlewares.ProductCreateMiddleware(), middlewares.AdminMiddleware(), handlers.CreateProduct)
		productRoutes.PUT("/:id", middlewares.ValidateBody(&product.Update{}), middlewares.AdminMiddleware(), handlers.UpdateProduct)
	}

	categoryRoutes := r.Group(baseURL + "/categories")
	{
		categoryRoutes.GET("", handlers.GetCategories)
	}

	return r
}
