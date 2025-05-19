package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/ClothShop/product-service/internal/utils"
)

func ValidateBody[T any](target *T) gin.HandlerFunc {
	log.Println("Validating body")
	return func(c *gin.Context) {
		if err := json.NewDecoder(c.Request.Body).Decode(target); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "incorrect json data"})
			log.Println("incorrect json data:", err)
			return
		}

		if err := utils.ValidateStruct(target); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println("Body validation error:", err)
			return
		}
		c.Set("validatedBody", target)
		c.Next()
	}
}
