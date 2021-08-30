package main

import (
	"html/template"
	"log"
	"net/http"
)

// Home page
func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")

}

func main() {
	// Build up the server:
	http.HandleFunc("/", homePage)
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

// Function for all templates, it has two parameters: http.ResponseWriter, and page
func renderTemplate(w http.ResponseWriter, page string) {
	// t for template
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil) 
	if err != nil {
		log.Println(err)
		return
	}
}