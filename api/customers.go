package api

import (
	"fmt"
	"gorm.io/gorm"

	"problem-company/models"
)

// Function to returns an array of up to 50 customers
func GetCustomers(db *gorm.DB) []models.Customer {
	var customers []models.Customer
	if result := db.Limit(50).Find(&customers); result.Error != nil {
        fmt.Println(result.Error)
    }

	return customers
}

// Function to return a customer by id
func GetCustomerById(db *gorm.DB, id string) models.Customer {
	var customer models.Customer
	if result := db.First(&customer, id); result.Error != nil {
        fmt.Println(result.Error)
    }

	return customer
}

// Function to create a customer
func CreateCustomer(db *gorm.DB, customer models.Customer) {
	if result := db.Create(&customer); result.Error != nil {
        fmt.Println(result.Error)
    }
}

// Function to update a customer
func UpdateCustomer(db *gorm.DB, customer models.Customer) {
	if result := db.Save(&customer); result.Error != nil {
        fmt.Println(result.Error)
    }
}