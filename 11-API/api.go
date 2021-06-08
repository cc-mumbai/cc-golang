package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

type Planet struct {
	Name string `json:"name"`
	Population string `json:"population"`
	Terrain string `json:"terrain"`
}

type Person struct {
	Name string `json:"name"`
	HomeWorldURL string `json:"homeworld"`
	HomeWorld Planet
}

// AllPeople is a collection of Person type
type AllPeople struct {
	People []Person `json:"results"`
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Getting people")
	res, err := http.Get(BaseURL + "people")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request Star Wars people")
	}
	fmt.Println(res)

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Fialed to parse request body")
	}
	var people AllPeople
	fmt.Println(string(bytes))

	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}
	fmt.Println(people)
}

func main() {
	http.HandleFunc("/people", getPeople)
	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}