package main

import "fmt"

func main() {

	// var myArray [5]int
	// var mySlice []int

	// myArray[0] = 1
	// mySlice[0] = 1

	// fmt.Println(myArray)
	// fmt.Println(mySlice)

	var raikage [5]int
	var hokage []int = make([]int, 5, 10)

	fmt.Println(raikage)
	fmt.Println(hokage)

	raikage[0] = 1
	hokage[0] = 1

	fmt.Println(raikage, len(raikage), cap(raikage))
	fmt.Println(hokage, len(hokage), cap(hokage))

	jutsuTypes := [5]string{"Ninjutsu", "Genjutsu", "Taijutsu", "Kekkei Genkai", "Senjutsu"}

	var splicedJutsu []string = jutsuTypes[1:3]

	fmt.Println(splicedJutsu, len(splicedJutsu), cap(splicedJutsu))

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1, slice2)
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))

	originalSlice := []int{1, 2, 3, 4}
	destination := make([]int, len(originalSlice))

	fmt.Println("Before Copy:", originalSlice, destination)

	mysteryValue := copy(destination, originalSlice)

	fmt.Println("After Copy:", originalSlice, destination, mysteryValue)
}