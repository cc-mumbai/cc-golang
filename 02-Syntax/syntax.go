package main

import (
	"fmt"
	"reflect"
)


func main() {
	fmt.Println("Sasuke!")

	// How to find type?   // reflect package
	var x = 4
	fmt.Println("4 is type:", reflect.TypeOf(x))

	// variables
	var jutsu string = "Ninjutsu"
	fmt.Println(jutsu)

	var clan = "Uchiha"
	fmt.Println(clan)

	var num int
	fmt.Println(num)

	var isHokage bool
	fmt.Println(isHokage)

	// shorthand variable syntax
	hero := "Itachi"
	fmt.Println(hero, "is type:", reflect.TypeOf(hero))
}