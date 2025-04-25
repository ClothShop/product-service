package middlewares

import (
	"github.com/ClothShop/product-service/internal/dto"
	"github.com/ClothShop/product-service/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProductUpdateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product dto.ProductUpdate
		if !utils.DecodeJSONRequest(c.Writer, c.Request, &product) {
			return
		}

		if err := utils.ValidateStruct(product); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Set("validatedProduct", product)
		c.Next()
	}
}
