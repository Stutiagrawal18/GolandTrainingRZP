package models

import (
	"gorm.io/gorm"
)

// Product represents a product in the retailer's inventory
type Product struct {
	gorm.Model
	ID          string  `gorm:"type:varchar(20);primaryKey" json:"id"`
	ProductName string  `gorm:"not null" json:"product_name"`
	Price       float64 `gorm:"not null" json:"price"`
	Quantity    int     `gorm:"not null" json:"quantity"`
}

// User represents a retailer user for authentication
type User struct {
	gorm.Model
	ID       string `gorm:"type:varchar(20);primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"` // Omit password from JSON
}

// Order represents a customer's order
type Order struct {
	gorm.Model
	ID          string          `gorm:"type:varchar(20);primaryKey" json:"id"`
	CustomerID  string          `gorm:"not null" json:"customer_id"`
	TotalPrice  float64         `gorm:"not null" json:"total_price"`
	Status      string          `gorm:"type:varchar(20);default:'placed'" json:"status"`
	LastOrderAt *gorm.DeletedAt `gorm:"index" json:"-"`
	OrderItems  []OrderItem     `gorm:"foreignKey:OrderID" json:"items"`
}

// OrderItem represents a single item in an order
type OrderItem struct {
	gorm.Model
	ID           string  `gorm:"type:varchar(20);primaryKey" json:"id"`
	OrderID      string  `gorm:"type:varchar(20);not null" json:"-"`
	ProductID    string  `gorm:"type:varchar(20);not null" json:"product_id"`
	Quantity     int     `gorm:"not null" json:"quantity"`
	PriceAtOrder float64 `gorm:"not null" json:"price_at_order"`
}
