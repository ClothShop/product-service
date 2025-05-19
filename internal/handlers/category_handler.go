package handlers

import (
	"github.com/ClothShop/product-service/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategories(c *gin.Context) {
	categories, err := services.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Error fetching categories",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, categories)
}
