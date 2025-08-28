package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DatabaseConfig holds all the database connection parameters.
type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

// GetDB returns a new database connection.
func GetDB(dbConfig DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}

// Add your own database configuration details here.
func MyDatabaseConfiguration() DatabaseConfig {
	return DatabaseConfig{
		User:     "root",              // e.g., "root"
		Password: "22bday@firstday",   // Your MySQL password
		Host:     "localhost",         // For a local connection
		Port:     "3306",              // Default MySQL port
		Database: "my_first_database", // The database you just created
	}
}
