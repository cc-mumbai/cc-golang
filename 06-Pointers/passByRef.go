package main

import (
	"fmt"
	"strings"
)

// Pass By Reference
func changeName(n *string) {
	*n = strings.ToUpper(*n)
}

func main() {
	name := "Rock Lee"
	changeName(&name)
	fmt.Println(name)
}