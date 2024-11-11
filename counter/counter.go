package counter

import (
	"fmt"
	"math/rand/v2"
)

func Testcall() {

	fmt.Println("hello from package")

}

type CounterSystem interface {
	Add(number int)
	Subtract(number int)
}

type Counter struct {
	sum int
}

func (c *Counter) Add(number int) {
	fmt.Println("add got called")
	c.sum = c.sum + number
	fmt.Println("New total value", c.sum)
}

func (c *Counter) Subtract(number int) {
	fmt.Println("Subtract got called")
	c.sum = c.sum - number
	fmt.Println("New total value", c.sum)
}

func (c *Counter) ExecRandom(number int) {
	var methods = []func(int){c.Add, c.Subtract}
	selectedIndex := rand.IntN(len(methods))
	methods[selectedIndex](number)
}

func ExecRandom2(counterSystem CounterSystem) {
	fmt.Println("test")

	roundsToPlay := rand.IntN(100)

	for roundNumber := range roundsToPlay {
	  fmt.Println("Round number ", roundNumber)
		methodToCall := rand.IntN(3)

		fmt.Println("Got number ", methodToCall)
	}

}
