// main.go
package main

import (
	config "first_api/Config"
	"log"
	// Make sure to use the correct path to your config package
)

func main() {
	// Call the MyDatabaseConfiguration function to get the connection details.
	dbConfig := config.MyDatabaseConfiguration()

	// Call the GetDB function with the configuration.
	db, err := config.GetDB(dbConfig)

	// Check for any errors. If 'err' is not nil, the connection failed.
	if err != nil {
		log.Fatal("Could not establish a database connection.")
	}

	// If the code reaches this point, the connection was successful.
	// You can now close the database connection.
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Could not get a database connection pool.")
	}
	defer sqlDB.Close()
}
