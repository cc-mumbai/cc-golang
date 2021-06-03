package main

import "fmt"

type Stats struct {
	Attack, Defence float64
}

var c = Stats{Attack: 10, Defence: 20}

func main() {
	NinjaStats := c   // What? No &? Go loophole hmm..
	NinjaStats.Attack = 15
	fmt.Println(NinjaStats)
}