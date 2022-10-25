package main

import (
	"problem-company/pkg/db"
	"problem-company/pkg/routes"
)

func main() {
	postgres.Connect()
	routes.StartRoutes()
}