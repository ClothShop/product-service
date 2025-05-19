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
	//r.SetTrustedProxies([]string{"127.0.0.1"})
	//
	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://localhost:8888"},
	//	AllowWildcard:    false,
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))

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
		//categoryRoutes.GET("/:id", handlers.GetCategory)
		//categoryRoutes.Use(middlewares.AuthMiddleware())
		//categoryRoutes.DELETE("/:id", middlewares.AdminMiddleware(), handlers.DeleteCategory)
		//
		//categoryRoutes.POST("", middlewares.ValidateBody(&category.Create{}), middlewares.AdminMiddleware(), handlers.CreateCategory)
		//categoryRoutes.PUT("", middlewares.ValidateBody(&category.Update{}), middlewares.AdminMiddleware(), handlers.UpdateCategory)
	}

	return r
}
