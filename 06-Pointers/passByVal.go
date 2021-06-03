package main

import (
	"fmt"
	"strings"
)

// Pass By Value
func changeName(n string) {
	n = strings.ToUpper(n)
}

func main() {
	name := "Rock Lee"
	changeName(name)
	fmt.Println(name)
}