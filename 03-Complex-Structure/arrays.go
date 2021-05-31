package main

import "fmt"

func main() {
	fmt.Println("Go arrays!")
	var scores [2]string
	scores[0] = "Naruto"
	scores[1] = "Boruto"
	fmt.Println(scores)

	fmt.Println("-----")
	kages := [5]float64{1, 2, 3, 4, 5}
	ninjaWars := [...]float64{1, 2, 3, 4}

	for _, value := range kages {
		fmt.Println(value)
	}

	fmt.Println("-----")
	for _, value := range ninjaWars {
		fmt.Println(value)
	}
}