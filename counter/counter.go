package counter

import (
	"fmt"
	"math/rand/v2"
)
 
type CounterSystem interface {
	Add(number int)
	Subtract(number int)
	Multiply(number int)
	Divide(number int)
	ExecRandom()
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

func (c *Counter) Multiply(number int){
  fmt.Printf("Will multiply by %d \n", number)
  c.sum = c.sum * number
  fmt.Println("New total value", c.sum)
}
  
func(c *Counter) Divide(number int){
  fmt.Println("Divide Called")
  if number == 0 {
    fmt.Println("Would divide by zero, return instead")
    return
  }
  c.sum = c.sum * number
  fmt.Println("New total value", c.sum) 
}

func (c *Counter) ExecRandom() {
		methodToCall := rand.IntN(4)
		//fmt.Println("Got number ", methodToCall)
		calcNumber := rand.IntN(10)
		switch methodToCall {
		case 0:
			fmt.Println("Will call add method!")
			c.Add(calcNumber)
		case 1:
			fmt.Println("Will call subtract method!")
			c.Subtract(calcNumber)
		case 2:
		  fmt.Println("Will call multiply method!")
		  c.Multiply(calcNumber)
		case 3: 
		  fmt.Println("Will call divide method!")
		  c.Divide(calcNumber)
		}
}

func NewCounter () *Counter {
  return new(Counter)
}
