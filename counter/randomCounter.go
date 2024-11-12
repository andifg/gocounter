package counter

import (
	"fmt"
	"math/rand/v2"
)

type RandomCounter struct {
	counter CounterSystem
}

func (r *RandomCounter) setCounter(counter CounterSystem) {
	r.counter = counter
}

func (r *RandomCounter) PlayRandomRounds() {

	if r.counter == nil {
		fmt.Println("ACHTUNG")
	}

	roundsToPlay := rand.IntN(100)
	for roundNumber := range roundsToPlay {
		fmt.Printf("Round: %d \n", roundNumber)
		r.counter.ExecRandom()
	}
}

func NewRandomCounter(counter CounterSystem) *RandomCounter {
	r := new(RandomCounter)
	r.setCounter(counter)
	return r
}
