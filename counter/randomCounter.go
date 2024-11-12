package counter

import (
	//"fmt"
	"math/rand/v2"
	log "github.com/sirupsen/logrus"
)

type RandomCounter struct {
	counter CounterSystem
}

type CounterSystem interface {
	Add(number int)
	Subtract(number int)
	Multiply(number int)
	Divide(number int)
	ExecRandom()
}


func (r *RandomCounter) setCounter(counter CounterSystem) {
	r.counter = counter
}

func (r *RandomCounter) PlayRandomRounds() {
	if r.counter == nil {
		log.Panic("ACHTUNG")
	}
	roundsToPlay := rand.IntN(100)
	for roundNumber := range roundsToPlay {
		log.Info("Round: ", roundNumber)
		r.counter.ExecRandom()
	}
}

func NewRandomCounter(counter CounterSystem) *RandomCounter {
	r := new(RandomCounter)
	r.setCounter(counter)
	return r
}
