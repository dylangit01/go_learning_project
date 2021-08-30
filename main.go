package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
)

// Home page
func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")

}

func playRound(w http.ResponseWriter, r *http.Request) {
	result := rps.PlayRound(1)

	out, err := json.MarshalIndent(result, "", "   ")

	if err != nil {
		log.Panicln(err)
		return
	}
	w.Header().Set("Content/Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)

	// Build up the server
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