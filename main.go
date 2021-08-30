package main

import (
	"fmt"
	"log"
	"net/http"
)

// Home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func main() {

	// Build up the server:
	http.HandleFunc("/", homePage)
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}