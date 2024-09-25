package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler function for the root route "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<html>
			<head>
				<link rel="stylesheet" href="/static/style.css">
			</head>
			<body>
				<h1>Welcome to the Go Web Server!</h1>
			</body>
		</html>
	`)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}

// Start the server
func main() {
	// Serve the static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define the route and corresponding handler function
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Start the HTTP server
	log.Println("Starting server on: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}

}
