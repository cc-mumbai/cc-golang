package main

import "fmt"

func main() {
	jutsu := "Taijutsu"

	for _, value := range jutsu {     // _ is a placeholder for a variable
		fmt.Println(string(value))
	}

	// print letters at odd positions
	for index, value := range jutsu {
		if index % 2 == 0 {
			fmt.Println(string(value))
		}
	}
}