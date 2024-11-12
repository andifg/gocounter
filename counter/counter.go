package counter

import (
	//"fmt"
	"math/rand/v2"
	log "github.com/sirupsen/logrus"
)

type Counter struct {
	sum int
}

func (c *Counter) Add(number int) {
	log.Debug("add got called")
	c.sum = c.sum + number
	log.Debug("New total value", c.sum)
}

func (c *Counter) Subtract(number int) {
	log.Debug("Subtract got called")
	c.sum = c.sum - number
	log.Debug("New total value", c.sum)
}

func (c *Counter) Multiply(number int){
  log.Debug("Will multiply by %d \n", number)
  c.sum = c.sum * number
  log.Debug("New total value", c.sum)
}
  
func(c *Counter) Divide(number int){
  log.Debug("Divide Called")
  if number == 0 {
    log.Debug("Would divide by zero, return instead")
    return
  }
  c.sum = c.sum * number
  log.Debug("New total value", c.sum) 
}

func (c *Counter) ExecRandom() {
		methodToCall := rand.IntN(4)
		calcNumber := rand.IntN(10)
		switch methodToCall {
		case 0:
			log.Trace("Will call add method!")
			c.Add(calcNumber)
		case 1:
			log.Trace("Will call subtract method!")
			c.Subtract(calcNumber)
		case 2:
		  log.Trace("Will call multiply method!")
		  c.Multiply(calcNumber)
		case 3: 
		  log.Trace("Will call divide method!")
		  c.Divide(calcNumber)
		}
}

func NewCounter () *Counter {
  return new(Counter)
}
