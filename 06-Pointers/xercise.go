/*
1. Define an instance of the Ninja struct
2. Write a function called updateEmail that takes in a *Ninja type
3. Update the Ninja's email to something new
4. Call updateEmail() from main() and verify the updated email has persisted
*/

package main

import "fmt"

type Ninja struct {
	ID int
	FirstName, LastName, Email string
}

var u = Ninja{
	ID: 1,
	FirstName: "Itachi",
	LastName:  "Uchiha",
	Email:     "itachi@akatsuki.com",
}

func updateEmail(u *Ninja) {
	u.Email = "itachi@konoha.com"
	fmt.Println("Updated Email:", u.Email)
}

func main() {
	fmt.Println("Pointers Exercise!")
	updateEmail(&u)
	fmt.Println("Updated Ninja:", u)
}