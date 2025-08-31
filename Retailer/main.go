package main

import (
	"log"
	"os"

	"Retailer/database"
	"Retailer/handlers"
	"Retailer/middleware"
	"Retailer/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Please ensure it exists and is configured.")
	}

	// Connect to the database
	database.ConnectDB()

	// Run database migrations to create the tables
	database.DB.AutoMigrate(&models.Product{}, &models.Order{}, &models.OrderItem{}, &models.User{})

	var count int64
	database.DB.Model(&models.User{}).Count(&count)
	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		defaultUser := models.User{
			Username: "retailer",
			Password: string(hashedPassword),
			ID:       "1",
		}
		database.DB.Create(&defaultUser)
		log.Println("Default retailer user created.")
	}
	// Initialize Gin router
	router := gin.Default()

	// Middleware for panic recovery
	router.Use(middleware.PanicRecovery())

	// Public routes (accessible by customers)
	router.GET("/products", handlers.GetProducts)
	router.GET("/product/:id", handlers.GetProduct)
	router.POST("/order", handlers.PlaceOrder)
	router.GET("/order/:id", handlers.GetOrder)
	router.GET("/customer/orders/:customer_id", handlers.GetCustomerOrderHistory)

	// Authentication route for retailers
	router.POST("/login", handlers.Login)

	// Protected routes (for retailers, requires JWT authentication)
	protected := router.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.POST("/product", handlers.CreateProduct)
		protected.PATCH("/product/:id", handlers.UpdateProduct)
		protected.GET("/business/orders", handlers.GetBusinessOrderHistory)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s...", port)
	log.Fatal(router.Run(":" + port))
}
