package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/global-ecommerce/internal/database"
	"github.com/yourusername/global-ecommerce/internal/models"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := database.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// Implement GetProduct, UpdateProduct, and DeleteProduct similarly
