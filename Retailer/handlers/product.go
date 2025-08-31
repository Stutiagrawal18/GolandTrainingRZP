package handlers

import (
	"Retailer/database"
	"Retailer/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	//"strconv"
	"sync"
)

// productMutex is used to prevent race conditions on product quantity updates.
var productMutex = &sync.Mutex{}

// CreateProduct godoc
// @Summary Add a new product
// @Description Adds a new product to the retailer's inventory.
// @Accept json
// @Produce json
// @Param product body models.Product true "Product object to be added"
// @Success 201 {object} models.Product
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /product [post]
func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct.ID = "PROD" + uuid.New().String()[:8]

	result := database.DB.Create(&newProduct)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": newProduct.ID, "product_name": newProduct.ProductName, "price": newProduct.Price, "quantity": newProduct.Quantity, "message": "product successfully added"})
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Updates the price and/or quantity of a product.
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param updates body models.Product true "Fields to update"
// @Success 200 {object} models.Product
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /product/{id} [patch]
func UpdateProduct(c *gin.Context) {
	productMutex.Lock()
	defer productMutex.Unlock()

	productID := c.Param("id")
	var product models.Product

	// Find the product
	if err := database.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var updates struct {
		Price    *float64 `json:"price"`
		Quantity *int     `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the product fields if they are provided
	if updates.Price != nil {
		product.Price = *updates.Price
	}
	if updates.Quantity != nil {
		product.Quantity = *updates.Quantity
	}

	// Save the updated product
	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProduct godoc
// @Summary Get a single product by ID
// @Description Retrieves a product's details.
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} gin.H
// @Router /product/{id} [get]
func GetProduct(c *gin.Context) {
	productID := c.Param("id")
	var product models.Product

	if err := database.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProducts godoc
// @Summary Get all products
// @Description Retrieves a list of all products in the catalog.
// @Produce json
// @Success 200 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
