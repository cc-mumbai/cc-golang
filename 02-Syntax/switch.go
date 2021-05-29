package main

import "fmt"

func main() {

	fmt.Println("Go switch")

	var clan string

	switch clan {
	case "Naruto":
		fmt.Println("Uzumaki")
	case "Itachi", "Obito":   // kinda like OR
		fmt.Println("Uchiha")
	case "Tsunade":
		fmt.Println("Senju")
	default:
		fmt.Println("That's it, Lee! Let the power of youth explode!!")
	}

	
	i := 11

	switch {
	case i > 10:
		fmt.Println("Greater than 10")
	case i < 10:
		fmt.Println("Less than 10")
	default:
		fmt.Println("Is 10")
	}


	var j int = 9

	switch {
	case j != 10:
		fmt.Println("Does not equal 10")
		fallthrough
	case j < 10:
		fmt.Println("Less than 10")
	case j > 10:
		fmt.Println("Greater than 10")
	default:
		fmt.Println("Is 10")
	}

}