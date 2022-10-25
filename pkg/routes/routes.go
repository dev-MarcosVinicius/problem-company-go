package routes

import (
    "io/ioutil"
    "encoding/json"
    "fmt"
    "net/http"
	"github.com/julienschmidt/httprouter"

	"problem-company/pkg/db"
	"problem-company/pkg/models"
	"problem-company/pkg/lib"
)

// Function to update a customer by id
func updateCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        fmt.Println("Error on read body.")
    }
    
    var updatedCustomer models.Customer
    json.Unmarshal(body, &updatedCustomer)

    var customer models.Customer

    customer = postgres.GetCustomerById(id)

    customer.First_Name = updatedCustomer.First_Name
    customer.Last_Name = updatedCustomer.Last_Name
    customer.Email = updatedCustomer.Email
    customer.Password = password.HashPassword(updatedCustomer.Email)

    postgres.UpdateCustomer(customer)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Updated")
}

// Function to create a customer
func createCustomer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        fmt.Println("Error on read body.")
    }
    
    var customer models.Customer
    json.Unmarshal(body, &customer)

    customer.Password, _ = password.HashPassword(customer.Password)
    
    postgres.CreateCustomer(customer)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}

// Function to return a customer by id
func getCustomerById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    customer := postgres.GetCustomerById(id)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(customer)
}

// Function to returns an array of up to 50 customers
func getCustomers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    customers := postgres.GetCustomers()

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(customers)
}

func StartRoutes() {
	router := httprouter.New()
    router.GET("/customers", getCustomers)
    router.GET("/customers/:id", getCustomerById)
    router.POST("/customers", createCustomer)
    router.PUT("/customers/:id", updateCustomer)

    fmt.Println("Running API on port: 1122")

    err := http.ListenAndServe("localhost:1122", router)
	if err != nil {
        panic(err)
	}
}