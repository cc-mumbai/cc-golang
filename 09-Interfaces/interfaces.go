package main

import "fmt"

// Describer interface describes the struct being passed
type Describer interface {
	describe() string
}

// User is a single user type
type User struct {
	ID int
	FirstName, LastName, Email string
}

// Group is a group of Users
type Group struct {
	role string
	users []User
	newestUser User
	spaceAvailable bool
}

// These two structs have different implementations of the `describe()` method.
func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

func (g *Group) describe() string {
	desc := fmt.Sprintf("The %s user group has %d users. Newest user: %s, Accepting New Users: %t", g.role, len(g.users), g.newestUser.FirstName, g.spaceAvailable)
	return desc
}

// DoTheDescribing describes whatever struct is passed
func DoTheDescribing(d Describer) string {
	return d.describe()
}

func main() {
	u1 := User{ID: 1, FirstName: "Jon", LastName: "Snow", Email: "jon@snow.com"}
	u2 := User{ID: 2, FirstName: "Tyrion", LastName: "Lannister", Email: "tyrion@lannister.com"}
	g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: true}
	
	describeUser := u1.describe()
	describeGroup := g.describe()

	fmt.Println(describeUser)
	fmt.Println(describeGroup)

	// Same task but w/ interface
	userDescInterface := DoTheDescribing(&u1)
	groupDescInterface := DoTheDescribing(&g)
	
	fmt.Println(userDescInterface)
	fmt.Println(groupDescInterface)
}