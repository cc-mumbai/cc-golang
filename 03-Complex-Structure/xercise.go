package main

import "fmt"

// Exercise 1
// 1. Instantiate an array of scores
//    The array should have at least 5 elements of type float64
// 2. Write a function that calculates and returns the average score(also a float)
//    Use the range keyword

// Exercise 2
// 1. Define a map that contains a set of pet names
// 	  and their corresponding breed. i.e.: "fido": "golden"
// 2. Write a function that takes a string argument and returns a boolean indicating
//    whether or not that key exists in your map of pets.

// Exercise 3
// 1. Instantiate a slice that has an initial value of a collection of groceries.
// 2. Write a function that takes one or more groceries as strings and
//    appends them to the slice, printing out the resulting list of groceries.

// Exercise 1 Solution
func average (scores [5]float64) float64 {
	total := 0.0
	for _, num := range scores {
		total += num
	}
	return total/float64(len(scores))
}

// Exercise 2 Solution
var woofMap map[string]string = map[string]string{
	"Tucker": "Golden",
	"Gatsby": "Corgi",
	"Ghost": "Husky",
	"Buzo": "German",
	"Maxine": "Corgi",
}

func isButtFluffy(dogName string) bool {
	_, ok := woofMap[dogName]
	return ok
}

// Exercise 3 Solution
var initGrocery = []string{"eggs", "milk", "beans", "bread"}

func addGrocery(newGroceries ...string) []string {
	foods := initGrocery
	for _, g := range newGroceries {
		foods = append(foods, g)
	}
	return foods
}

func main() {
	fmt.Println("Exercise 1 Solution")
	scores := [5]float64{1, 2, 3, 4, 5}
	fmt.Println(average(scores))
	fmt.Println("Exercise 2 Solution")
	dog := "Gatsby" 
	fluffyButt := isButtFluffy(dog)
	fmt.Println(fluffyButt)
	fmt.Println("Exercise 3 Solution")
	groceryList := addGrocery("juice", "mangoes")
	fmt.Println(groceryList)
}