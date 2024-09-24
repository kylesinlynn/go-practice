package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler function for the root route "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Web Server!")
}

// Start the server
func main() {
	// Define the route and corresponding handler function
	http.HandleFunc("/", homeHandler)

	// Start the HTTP server
	log.Println("Starting server on: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}

}
