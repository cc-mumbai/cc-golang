package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Go Home!")
}

func todos(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("todos.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error: ", err)
	}
	err = t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", todos)
	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}