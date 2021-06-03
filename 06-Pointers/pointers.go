package main

import "fmt"

func main() {
	fmt.Println("Go Pointers")
	var name string
	var namePtr *string

	name = "Kisame"
	namePtr = &name
	var nameVal = *namePtr

	fmt.Println("Name:", name)
	fmt.Println("NamePtr:", namePtr)
	fmt.Println("NameVal:", nameVal)
}