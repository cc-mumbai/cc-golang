package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func isLevelTen(num int) error {
	if num < 10 {
		return errors.New("That's what she said")
	}
	return nil
}

func openFile() error {
	f, err := os.Open("secret.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func main() {
	fmt.Println("Go Errors!")

	num := 9
	// err := isLevelTen(num)
	if err := isLevelTen(num); err != nil {
		fmt.Println(fmt.Errorf("%d is not greater than 10", num))
		// panic(err)   // EOD
		// fmt.Println("---")   // Code Unreachable after panic()
		// log.Fatalln(err)   // Keeps logs for reference
	}
	// fmt.Println(err)   // err is only scoped to the above if block

	// err2 := openFile()
	if err := openFile(); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
		log.Fatalln(err)
	}

}