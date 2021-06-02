package main

import "fmt"

// Struct
type Ninja struct {
	Level int
	FirstName, LastName, Email string
}

// Custom Type
type Anbu struct {
	role string
	ninjas []Ninja
	newNinja Ninja
	spaceAvailable bool
}

func describeNinja(n Ninja) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", n.FirstName, n.LastName, n.Email)
	return desc
}

func main() {
	genin := Ninja{
		Level: 1, 
		FirstName: "Naruto", 
		LastName: "Uzumaki", 
		Email: "naruto@konoha.com",
	}
	fmt.Println(genin.Email)
	fmt.Println(genin)
	fmt.Println(describeNinja(genin))
}