package main

import (
    "os"
	"fmt"
    "net/http"
	"github.com/julienschmidt/httprouter"
    "io/ioutil"
    "encoding/json"
	"gorm.io/gorm"

	"problem-company/pkg"
	"problem-company/models"
	"problem-company/api"
)

// DBConnection is the class of the connection to database
type DBConnection struct {
	DB *gorm.DB
}

func main() {
	db := api.Connect("user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable")

	port := func() string { if os.Getenv("PORT") != "" { return os.Getenv("PORT") } else { return "1122" } }()
	router := httprouter.New()

	routes := &DBConnection{DB: db}

    router.GET("/customers", routes.GetCustomers)
    router.GET("/customers/:id", routes.GetCustomerById)
    router.POST("/customers", routes.CreateCustomer)
    router.PUT("/customers/:id", routes.UpdateCustomer)

    fmt.Println("Running API on port: " + port)

    err := http.ListenAndServe("localhost:" + port, router)
	if err != nil {
        panic(err)
	}
}


// Function to update a customer by id
func (postgres *DBConnection) UpdateCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        fmt.Println("Error on read body.")
    }
    
    var updatedCustomer models.Customer
    json.Unmarshal(body, &updatedCustomer)

    var customer models.Customer

    customer = api.GetCustomerById(postgres.DB, id)

    customer.First_Name = updatedCustomer.First_Name
    customer.Last_Name = updatedCustomer.Last_Name
    customer.Email = updatedCustomer.Email
    customer.Password, _ = password.HashPassword(updatedCustomer.Email)

    api.UpdateCustomer(postgres.DB, customer)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Updated")
}

// Function to create a customer
func (postgres *DBConnection) CreateCustomer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        fmt.Println("Error on read body.")
    }
    
    var customer models.Customer
    json.Unmarshal(body, &customer)

    customer.Password, _ = password.HashPassword(customer.Password)
    
    api.CreateCustomer(postgres.DB, customer)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}

// Function to return a customer by id
func (postgres *DBConnection) GetCustomerById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    customer := api.GetCustomerById(postgres.DB, id)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(customer)
}

// Function to returns an array of up to 50 customers
func (postgres *DBConnection) GetCustomers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    customers := api.GetCustomers(postgres.DB)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(customers)
}