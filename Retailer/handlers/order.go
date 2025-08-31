package handlers

import (
	"Retailer/database"
	"Retailer/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// The cool-down period for a customer's next order
const coolDownPeriod = 5 * time.Minute

// PlaceOrder godoc
// @Summary Place a new order
// @Description Creates a new order for one or more products.
// @Accept json
// @Produce json
// @Param order body object{customer_id=string,items=[]object{product_id=string,quantity=int}} true "Order details"
// @Success 201 {object} models.Order
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /order [post]
func PlaceOrder(c *gin.Context) {
	var requestBody struct {
		CustomerID string `json:"customer_id"`
		Items      []struct {
			ProductID string `json:"product_id"`
			Quantity  int    `json:"quantity"`
		} `json:"items"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(requestBody.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order must contain at least one item"})
		return
	}

	// Check for cooldown period
	var lastOrder models.Order
	if err := database.DB.Where("customer_id = ? AND status = ?", requestBody.CustomerID, "processed").Order("created_at DESC").First(&lastOrder).Error; err == nil {
		if time.Since(lastOrder.CreatedAt) < coolDownPeriod {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Please wait for the 5-minute cooldown period before placing another order"})
			return
		}
	}

	// Create a new order with a unique ID and a "placed" status
	newOrder := models.Order{
		ID:         "ORD" + uuid.New().String()[:8],
		CustomerID: requestBody.CustomerID,
		Status:     "placed",
	}

	// Use a transaction for atomicity
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&newOrder).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
		return
	}

	var totalPrice float64
	var orderItems []models.OrderItem
	var failed bool

	for _, item := range requestBody.Items {
		var product models.Product

		// Concurrency handling with mutex
		productMutex.Lock()
		if err := tx.Where("id = ?", item.ProductID).First(&product).Error; err != nil {
			failed = true
			productMutex.Unlock()
			continue
		}

		if product.Quantity < item.Quantity {
			failed = true
			productMutex.Unlock()
			continue
		}

		product.Quantity -= item.Quantity
		if err := tx.Save(&product).Error; err != nil {
			failed = true
			productMutex.Unlock()
			continue
		}

		productMutex.Unlock()

		orderItem := models.OrderItem{
			ID:           "ORDITEM" + uuid.New().String()[:8],
			OrderID:      newOrder.ID,
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			PriceAtOrder: product.Price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			failed = true
			continue
		}
		orderItems = append(orderItems, orderItem)
		totalPrice += product.Price * float64(item.Quantity)
	}

	if failed {
		tx.Rollback()
		newOrder.Status = "failed"
		database.DB.Save(&newOrder)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order could not be processed due to insufficient quantity or invalid product", "order_id": newOrder.ID, "status": newOrder.Status})
		return
	}

	newOrder.Status = "processed"
	newOrder.TotalPrice = totalPrice
	tx.Save(&newOrder)
	tx.Commit()

	newOrder.OrderItems = orderItems
	c.JSON(http.StatusCreated, newOrder)
}

// GetOrder godoc
// @Summary Get an order by ID
// @Description Retrieves a specific order's details.
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.Order
// @Failure 404 {object} gin.H
// @Router /order/{id} [get]
func GetOrder(c *gin.Context) {
	orderID := c.Param("id")
	var order models.Order

	if err := database.DB.Preload("OrderItems").Where("id = ?", orderID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetCustomerOrderHistory godoc
// @Summary Get customer's order history
// @Description Retrieves all past orders for a specific customer.
// @Produce json
// @Param customer_id path string true "Customer ID"
// @Success 200 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customer/orders/{customer_id} [get]
func GetCustomerOrderHistory(c *gin.Context) {
	customerID := c.Param("customer_id")
	var orders []models.Order
	if err := database.DB.Preload("OrderItems").Where("customer_id = ?", customerID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customer orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// GetBusinessOrderHistory godoc
// @Summary Get all business transactions
// @Description Retrieves all processed orders for the retailer.
// @Produce json
// @Success 200 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /business/orders [get]
func GetBusinessOrderHistory(c *gin.Context) {
	var orders []models.Order
	if err := database.DB.Preload("OrderItems").Where("status = ?", "processed").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch business orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": orders})
}
