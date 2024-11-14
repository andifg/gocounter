package counter

import (
	//"fmt"
	//"errors"
	log "github.com/sirupsen/logrus"
	"math/rand/v2"
)

type Counter struct {
	sum int
}

func (c *Counter) Add(number int) CounterOperationReport {
	log.Debug("add got called")
	c.sum = c.sum + number
	log.Debug("New total value", c.sum)
	return CounterOperationReport{
		mathOperation: Add,
		currentSum:    c.sum,
	}
}

func (c *Counter) Subtract(number int) CounterOperationReport {
	log.Debug("Subtract got called")
	c.sum = c.sum - number
	log.Debug("New total value", c.sum)
	return CounterOperationReport{
		mathOperation: Subtract,
		currentSum:    c.sum,
	}

}

func (c *Counter) Multiply(number int) CounterOperationReport {
	log.Debug("Will multiply by ", number)
	c.sum = c.sum * number
	log.Debug("New total value", c.sum)
	return CounterOperationReport{
		mathOperation: Multiply,
		currentSum:    c.sum,
	}
}

func (c *Counter) Divide(number int) CounterOperationReport {
	log.Debug("Divide Called")
	if number == 0 {
		log.Debug("Would divide by zero, return instead")
		return CounterOperationReport{
			mathOperation: Divide,
			currentSum:    c.sum,
		}
	}
	c.sum = c.sum * number
	log.Debug("New total value", c.sum)
	return CounterOperationReport{
		mathOperation: Divide,
		currentSum:    c.sum,
	}
}

func (c *Counter) ExecRandom() CounterOperationReport {
	methodToCall := rand.IntN(4)
	calcNumber := rand.IntN(10)
	switch methodToCall {
	case 0:
		log.Trace("Will call add method!")
		return c.Add(calcNumber)
	case 1:
		log.Trace("Will call subtract method!")
		return c.Subtract(calcNumber)
	case 2:
		log.Trace("Will call multiply method!")
		return c.Multiply(calcNumber)
	case 3:
		log.Trace("Will call divide method!")
		return c.Divide(calcNumber)
	}
	panic("Impossible case")
}

func NewCounter() *Counter {
	return new(Counter)
}
