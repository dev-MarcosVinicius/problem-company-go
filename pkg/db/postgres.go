package postgres

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"problem-company/pkg/models"
)

// Model Struct of Customer to Database
type Customer struct {
	gorm.Model
	First_Name  string
	Last_Name string
	Email string
	Password string
}

// Constant to save database connection
var db *gorm.DB

func Connect() {
	// Open the connection with Database
	connection, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=postgres password=postgres dbname=postgres host=::1 port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	  }), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected on Database!")

	// Migrate the schema of Customer
	connection.AutoMigrate(&Customer{})

	db = connection
}

// Function to returns an array of up to 50 customers
func GetCustomers() []models.Customer {
	var customers []models.Customer
	if result := db.Limit(50).Find(&customers); result.Error != nil {
        fmt.Println(result.Error)
    }

	return customers
}

// Function to return a customer by id
func GetCustomerById(id string) models.Customer {
	var customer models.Customer
	if result := db.First(&customer, id); result.Error != nil {
        fmt.Println(result.Error)
    }

	return customer
}

// Function to create a customer
func CreateCustomer(customer models.Customer) {
	if result := db.Create(&customer); result.Error != nil {
        fmt.Println(result.Error)
    }
}