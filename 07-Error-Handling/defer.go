package main

import "fmt"

// LIFO?
func whatIsDefer() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	fmt.Println("Third")
}

func main() {
	whatIsDefer()
}