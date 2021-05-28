package main

import (
	"errors"
	"fmt"
)

func someFunction() error {
	return errors.New("some error")
}

func main() {

	var shurikens = 9

	if shurikens > 10 {
		fmt.Println(shurikens)
	}

	if shurikens > 100 {
		fmt.Println("Greater than 100")
	} else if shurikens == 100 {
		fmt.Println("Equals 100")
	} else {
		fmt.Println("Less than 100")
	}

	err := someFunction()
	// => If this function returns a value, it will be an error of type Error

	if err != nil {
	  fmt.Println(err.Error())
	}

	if err := someFunction(); err != nil {
	  fmt.Println(err.Error())
	}

}