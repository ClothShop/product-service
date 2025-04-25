package middlewares

import (
	"encoding/json"
	"github.com/ClothShop/product-service/internal/dto"
	"github.com/ClothShop/product-service/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProductCreateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "❌ Failed to parse multipart form",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		productJson := c.Request.FormValue("product")
		var product dto.ProductCreate
		if err := json.Unmarshal([]byte(productJson), &product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "❌ Invalid product JSON",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		if err := utils.ValidateStruct(product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "❌ Validation failed",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("validatedProduct", product)
		c.Next()
	}
}
