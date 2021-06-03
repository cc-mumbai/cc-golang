package main

import "fmt"

func saveMe() {
	if r := recover(); r != nil {
		fmt.Println("Ryan started the fire")
		fmt.Println(r)   // PANIC!
	}
}

func whatIsRecover() {
	defer saveMe()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			panic("PANIC!")   // panic => defer => recover
		}
	}
}

func main() {
	whatIsRecover()
}