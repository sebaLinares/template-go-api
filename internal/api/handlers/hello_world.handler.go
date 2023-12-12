package handlers

import (
	"fmt"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, World!" as the response
	fmt.Fprint(w, "Hello, World!")
}
