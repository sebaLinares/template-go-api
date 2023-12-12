package main

import (
	"fmt"
	"net/http"

	"mygoapi/internal/api/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Create a Gorilla Mux router
	router := mux.NewRouter()

	// Define a route for the /hello-world endpoint
	routes.SetupHelloWorldRoutes(router)

	// Start the HTTP server on port 8080
	http.Handle("/", router)
	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
