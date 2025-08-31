package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	//"time"

	"Retailer/database"
	"Retailer/handlers"
	"Retailer/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupTestDB initializes an in-memory SQLite database for testing.
func SetupTestDB(t *testing.T) {
	var err error
	database.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schemas
	database.DB.AutoMigrate(&models.Product{}, &models.Order{}, &models.OrderItem{}, &models.User{})
}

// ClearTestDB clears the tables after each test
func ClearTestDB() {
	database.DB.Exec("DELETE FROM products")
	database.DB.Exec("DELETE FROM orders")
	database.DB.Exec("DELETE FROM order_items")
}

func TestGetProducts(t *testing.T) {
	SetupTestDB(t)
	defer ClearTestDB()

	// Create some dummy data
	database.DB.Create(&models.Product{ID: "PROD1", ProductName: "test-product-1", Price: 10.0, Quantity: 5})
	database.DB.Create(&models.Product{ID: "PROD2", ProductName: "test-product-2", Price: 20.0, Quantity: 10})

	router := gin.Default()
	router.GET("/products", handlers.GetProducts)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string][]models.Product
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Len(t, response["products"], 2)
}

func TestCreateProduct(t *testing.T) {
	SetupTestDB(t)
	defer ClearTestDB()

	router := gin.Default()
	router.POST("/product", handlers.CreateProduct)

	body := gin.H{
		"product_name": "test-product-3",
		"price":        30.0,
		"quantity":     15,
	}
	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Contains(t, response, "message")
	assert.Equal(t, "product successfully added", response["message"])

	// Check if the product exists in the DB
	var product models.Product
	database.DB.First(&product, "product_name = ?", "test-product-3")
	assert.Equal(t, "test-product-3", product.ProductName)
}

func TestPlaceOrder(t *testing.T) {
	SetupTestDB(t)
	defer ClearTestDB()

	// Seed product
	productID := "PROD" + uuid.New().String()[:8]
	database.DB.Create(&models.Product{ID: productID, ProductName: "test-item", Price: 50, Quantity: 10})

	router := gin.Default()
	router.POST("/order", handlers.PlaceOrder)

	body := gin.H{
		"customer_id": "CST1",
		"items": []gin.H{
			{
				"product_id": productID,
				"quantity":   2,
			},
		},
	}
	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/order", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Check if the order was created
	var order models.Order
	database.DB.Preload("OrderItems").First(&order, "customer_id = ?", "CST1")
	assert.Equal(t, float64(100), order.TotalPrice)
	assert.Equal(t, "processed", order.Status)

	// Check if product quantity was updated
	var product models.Product
	database.DB.First(&product, "id = ?", productID)
	assert.Equal(t, 8, product.Quantity)
}
