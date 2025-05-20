package handlers

import (
	"github.com/ClothShop/product-service/internal/dto/product"
	"github.com/ClothShop/product-service/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProduct(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	productDto, err := services.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {
	products, err := services.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching products",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	bodyAny, exists := c.Get("validatedBody")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid product context"})
		return
	}
	productCreate := bodyAny.(product.Create)

	form, err := c.MultipartForm()
	if err == nil {
		productCreate.Images = form.File["images"]
	}

	productDto, err := services.CreateProduct(&productCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating product",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, productDto)
}

func UpdateProduct(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	productAny, exists := c.Get("validatedBody")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid product context"})
		return
	}
	productDto := productAny.(*product.Update)

	if err := services.UpdateProduct(id, productDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error updating product",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, productDto)
}

func DeleteProduct(c *gin.Context) {
	id, ok := parseIDParam(c)
	if !ok {
		return
	}
	if err := services.DeleteProduct(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
			"error":   err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func parseIDParam(c *gin.Context) (uint, bool) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return 0, false
	}
	return uint(id), true
}
