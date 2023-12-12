package routes

import (
	"mygoapi/internal/api/handlers"

	"github.com/gorilla/mux"
)

func SetupHelloWorldRoutes(router *mux.Router) {
	router.HandleFunc("/api/hello-world", handlers.HelloWorldHandler).
		Methods("GET")
}
