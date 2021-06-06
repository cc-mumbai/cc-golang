package main

import "fmt"

// Methods allows us to have same function name!
type Ninja struct {
	ID int
	FirstName, LastName, Email string
}

// Group represents a set of ninjas
type Group struct {
	role string
	ninjas []Ninja
	newestUser Ninja
	spaceAvailable bool
}

func (u *Ninja) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

func (g *Group) describe() string {
	if len(g.ninjas) > 2 {
		g.spaceAvailable = false
	}

	desc := fmt.Sprintf("This Ninja group has %d members. The newest ninja is %s %s. Accepting New Ninjas: %t", len(g.ninjas), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc

}

func main() {
	u := Ninja{ID: 1, FirstName: "Madara", LastName: "Uchiha", Email: "madara@epic.com"}

	u2 := Ninja{ID: 2, FirstName: "Minato", LastName: "Namikaze", Email: "minato@hakage.com"}

	u3 := Ninja{ID: 3, FirstName: "Might", LastName: "Guy", Email: "guy@badass.com"}

	g := Group{
		role: "admin",
		ninjas: []Ninja{u, u2, u3},
		newestUser: u3,
		spaceAvailable: true,
	}

	fmt.Println(g.describe())
	fmt.Println(u.describe())
	fmt.Println(g)
}