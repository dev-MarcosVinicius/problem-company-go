package api

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"problem-company/models"
)

func Connect(dsn string) *gorm.DB {
	// Open the connection with Database
	connection, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	  }), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected on Database!")

	// Migrate the schema of Customer
	connection.AutoMigrate(&models.Customer{})

	return connection
}