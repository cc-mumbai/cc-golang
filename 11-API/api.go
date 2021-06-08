package main

import (
	"fmt"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

func getPeople(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Getting people")
	res, err := http.Get(BaseURL + "people")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Failed to request Star Wars people")
	}
	fmt.Println(res)
}

func main() {
	http.HandleFunc("/people", getPeople)
	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}