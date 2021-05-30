package main

import "fmt"

// Exercise 1
// 1. Create a function called average that takes three separate arguments,
//    all of which are floats.
// 2. The function should return the average of the three arguments as a float.

// Exercise 2
// Refactor your code to use a variadic function
// that takes in an unknown number of arguments.
// HINT: To find the length of a collection, use len(xyzCollection)

// Exercise 1 Solution
func average(arg1, arg2, arg3 float64) float64 {
	total := arg1 + arg2 + arg3
	return total/3
}

// Exercise 2 Solution
func avg(nums ...float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total/float64(len(nums))
}

func main() {
	fmt.Println("Functions exercise!")
	fmt.Println(average(3, 4, 5))
	fmt.Println(avg(1, 2, 3, 4, 5))
}