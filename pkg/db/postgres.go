package postgres

import (
	"fmt"
    "os"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

// Model Struct of Customer to Database
type Customer struct {
	gorm.Model
	First_Name  string
	Last_Name string
	Email string
	Password string
}

func Connect() *gorm.DB {
	port := os.Getenv("PORT")
	dsn := func() string {	
		if port != "" {
			return "user=postgres password=postgres dbname=postgres host=::1 port=" + os.Getenv("PORT") + "sslmode=disable"
		} else {
			return "user=postgres password=postgres dbname=postgres host=::1 port=5432 sslmode=disable"
		}
	}
	
	// Open the connection with Database
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn(),
		PreferSimpleProtocol: true,
	  }), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected on Database!")

	// Migrate the schema of Customer
	db.AutoMigrate(&Customer{})

    return db
}