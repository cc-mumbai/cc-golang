package main

import "fmt"

func main() {

	// var secretScrolls map[int]string

	// secretScrolls[1] = "user1@gmail.com"
	// secretScrolls[2] = "user2@gmail.com"

	// fmt.Println(secretScrolls)


	// var secretScrolls2 map[int]string = make(map[int]string)
	secretScrolls2 := make(map[int]string)

	secretScrolls2[1] = "user1@gmail.com"
	secretScrolls2[2] = "user2@gmail.com"

	fmt.Println(secretScrolls2)
	fmt.Println("-----")


	secretScrolls3 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	fmt.Println(secretScrolls3)
	fmt.Println(secretScrolls3[1])
	fmt.Println("-----")


	secretScrolls4 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	fmt.Println(secretScrolls4)

	secretScrolls4[1] = "newUser1@gmail.com"

	fmt.Println(secretScrolls4)
	fmt.Println(secretScrolls4[3])
	fmt.Println("-----")


	secretScrolls5 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	email1, ok := secretScrolls5[1]
	fmt.Println("Email:", email1, "Present?", ok)

	email3, ok := secretScrolls5[3]
	fmt.Println("Email", email3, "Present?", ok)
	fmt.Println("-----")


	secretScrolls6 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	if email, ok := secretScrolls6[4]; ok {
		fmt.Println(email)
	} else {
		fmt.Println("I don't know what you want from me")
	}

	delete(secretScrolls6, 2)
	fmt.Println(secretScrolls6)
	fmt.Println("-----")


	secretScrolls7 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	for k, v := range secretScrolls7 {
		fmt.Printf("%s has an ID of %d.\n", v, k)
	}
}