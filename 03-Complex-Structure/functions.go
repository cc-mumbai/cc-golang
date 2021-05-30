package main

import "fmt"

func age(x int) int {
	fmt.Println(x)
	return x
}

func printAge(age1, age2 int) (ageOfSasuke, ageOfItachi int) {
	ageOfSasuke = age1
	ageOfItachi = age2
	return
}

func spreadAge(ages ...int) {
	for _, age := range ages {
		fmt.Println(age)
	}
}

func main() {
	fmt.Println("Hello functions!")
	age(21)
	fmt.Println("-----")
	x, y := printAge(10, 18)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("-----")
	spreadAge(12, 24, 36)
}