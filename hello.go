package main

import (
	"fmt"
	//"math/rand/v2"
	"playground/hello/counter"
)

func main() {
	fmt.Println("###############")
	counterInstance := counter.NewCounter()
	randomCounter := counter.NewRandomCounter(counterInstance)
	randomCounter.PlayRandomRounds()
}
