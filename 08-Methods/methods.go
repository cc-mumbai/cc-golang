package main

import "fmt"

type Ninja struct {
	ID int
	FirstName, LastName, Email string
}

// Function
func describeNinja(n Ninja) string {
	desc := fmt.Sprintf("Name: %s %s | Email: %s | ID: %d", n.FirstName, n.LastName, n.Email, n.ID)
	return desc
}

// Method
func (n *Ninja) describe() string {
	desc := fmt.Sprintf("ID: %d | Name: %s %s | Email: %s", n.ID, n.FirstName, n.LastName, n.Email)
	return desc
}

func main() {
	fmt.Println("Go methods!")
	genin := Ninja{ID: 1, FirstName: "Choji", LastName: "Akimichi", Email: "choji@food.com"}
	
	descFn := describeNinja(genin)
	fmt.Println(descFn)

	descMtd := genin.describe()
	fmt.Println(descMtd)
}