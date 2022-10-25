package routes

import (
    "encoding/json"
    "fmt"
    "net/http"
	"github.com/julienschmidt/httprouter"

	"problem-company/pkg/db"
)

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

    fmt.Println("Running API on port: 1122")

    err := http.ListenAndServe("localhost:1122", router)
	if err != nil {
        panic(err)
	}
}