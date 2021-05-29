package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Do-while loop does not exist in Golang😮")

	i := 1

	for i <= 5 {
		fmt.Println(i)
		// This will behave like a while loop?
		i += 1
	}


	var mySentence = "Akatsuki"

	for index, letter := range mySentence {
		fmt.Println("Index:", index, "Letter:", string(letter))
	}

}