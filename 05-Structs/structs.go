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
	if len(g.ninjas) > 2 {
		g.spaceAvailable = false
	}
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

	og_anbu := Ninja{
		Level: 9, 
		FirstName: "Itachi", 
		LastName: "Uchiha", 
		Email: "itachi@akatsuki.com",
	}

	tsukuyomi := Anbu{
		role: "admin",
		ninjas: []Ninja{genin, anbu, og_anbu},
		newNinja: og_anbu,
		spaceAvailable: true,
	}

	fmt.Println(genin.Email)
	fmt.Println(genin)
	fmt.Println(describeNinja(genin))
	fmt.Println("-----")
	fmt.Println(describeGroup(tsukuyomi))
	fmt.Println("-----")
	fmt.Println(tsukuyomi)   // Still True is shown hmm:( 
	// Modifying a variable permanently requires more behaviour:)
}