package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Handler function for the root route "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home.html")
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Home Page",
		Message: "Welcome to the home page",
	}
	tmpl.Execute(w, data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		fmt.Fprintf(w, "Hello, %s! Welcome to the submit page", username)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Start the server
func main() {
	// Serve the static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define the route and corresponding handler function
	http.Handle("/", loggingMiddleware(http.HandlerFunc(homeHandler)))
	http.Handle("/about", loggingMiddleware(http.HandlerFunc(aboutHandler)))
	http.Handle("/submit", loggingMiddleware(http.HandlerFunc(submitHandler)))

	// Start the HTTP server
	log.Println("Starting server on: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}

}
