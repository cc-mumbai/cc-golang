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

func describeGroup(g Anbu) string {
	desc := fmt.Sprintf("This Ninja group has %d members. The newest ninja is %s %s. Accepting new ninjas: %t", len(g.ninjas), g.newNinja.FirstName, g.newNinja.LastName, g.spaceAvailable)
	return desc
}

func main() {
	genin := Ninja{
		Level: 1, 
		FirstName: "Naruto", 
		LastName: "Uzumaki", 
		Email: "naruto@konoha.com",
	}

	anbu := Ninja{
		Level: 7, 
		FirstName: "Kakashi", 
		LastName: "Hatake", 
		Email: "kakashi@konoha.com",
	}

	tsukuyomi := Anbu{
		role: "admin",
		ninjas: []Ninja{genin, anbu},
		newNinja: anbu,
		spaceAvailable: true,
	}

	fmt.Println(genin.Email)
	fmt.Println(genin)
	fmt.Println(describeNinja(genin))
	fmt.Println(describeGroup(tsukuyomi))
}