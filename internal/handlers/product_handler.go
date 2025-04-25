package handlers

import (
	"github.com/ClothShop/product-service/internal/dto"
	"github.com/ClothShop/product-service/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Service *services.ProductService
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "❌ Invalid product ID",
		})
		return
	}

	product, err := h.Service.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "❌ Product not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.Service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Error fetching products",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
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

	if err := h.Service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Error creating product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productAny, exists := c.Get("validatedProduct")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Invalid product context",
		})
		return
	}
	product := productAny.(dto.ProductUpdate)

	if err := h.Service.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "❌ Error updating product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "❌ Invalid product ID",
		})
		return
	}

	if err := h.Service.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "❌ Product not found",
			"error":   err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
