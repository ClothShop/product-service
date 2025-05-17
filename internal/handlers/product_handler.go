package handlers

import (
	"github.com/ClothShop/product-service/internal/dto"
	"github.com/ClothShop/product-service/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "❌ Invalid product ID",
		})
		return
	}

	product, err := services.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "❌ Product not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetProducts(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Error fetching products",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	productAny, exists := c.Get("validatedProduct")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Invalid product context",
		})
		return
	}
	product := productAny.(dto.ProductCreate)

	form, err := c.MultipartForm()
	if err == nil {
		files := form.File["images"]
		product.Images = files
	}

	productDto, err := services.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Error creating product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, productDto)
}

func UpdateProduct(c *gin.Context) {
	productAny, exists := c.Get("validatedProduct")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Invalid product context",
		})
		return
	}
	product := productAny.(dto.ProductUpdate)

	if err := services.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "❌ Error updating product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "❌ Invalid product ID",
		})
		return
	}

	if err := services.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "❌ Product not found",
			"error":   err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
