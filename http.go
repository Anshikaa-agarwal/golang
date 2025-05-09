package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Root handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)     // GET, POST, etc.
	fmt.Println("URL Path:", r.URL.Path) // /hello
	fmt.Println("Headers:", r.Header)    // Map of headers
	fmt.Println("Query Param:", r.URL.Query().Get("id"))
}

// About handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the About Page.")
}

// JSON handler
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Hello from JSON route!",
		"status":  "success",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)           // Route: /
	http.HandleFunc("/about", aboutHandler) // Route: /about
	http.HandleFunc("/json", jsonHandler)   // Route: /json

	fmt.Println("Server is running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
