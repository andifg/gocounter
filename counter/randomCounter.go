package counter

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand/v2"
)

type RandomCounter struct {
	counter   CounterSystem
	statistic RandomCounterStatistic
}

type RandomCounterStatistic struct {
	TotalSum      int `json:"totalSum"`
	TotalRounds   int `json:"totalRounds"`
	AddCount      int `json:"addCount"`
	SubtractCount int `json:"subtractCount"`
	MultiplyCount int `json:"multiplyCount"`
	DivideCount   int `json:"divideCount"`
}

type MathOperation int

const (
	Add MathOperation = iota
	Subtract
	Multiply
	Divide
)

type CounterOperationReport struct {
	mathOperation MathOperation
	currentSum    int
}

type CounterSystem interface {
	Add(number int) CounterOperationReport
	Subtract(number int) CounterOperationReport
	Multiply(number int) CounterOperationReport
	Divide(number int) CounterOperationReport
	ExecRandom() CounterOperationReport
}

func (r *RandomCounter) setCounter(counter CounterSystem) {
	r.counter = counter
}

func (r *RandomCounter) PlayRandomRounds() {
	if r.counter == nil {
		log.Panic("ACHTUNG")
	}
	roundsToPlay := rand.IntN(100)
	log.Info("Will play ", roundsToPlay, " rounds")
	for roundNumber := range roundsToPlay {
		log.Debug("Round: ", roundNumber)
		r.executeRandomProcessResult()
	}
	r.printSummary()
}

func (r *RandomCounter) executeRandomProcessResult() {
	res := r.counter.ExecRandom()
	r.statistic.TotalRounds += 1
	r.statistic.TotalSum += res.currentSum
	switch res.mathOperation {
	case Add:
		r.statistic.AddCount += 1
	case Subtract:
		r.statistic.SubtractCount += 1
	case Multiply:
		r.statistic.MultiplyCount += 1
	case Divide:
		r.statistic.DivideCount += 1
	}
}

func (r *RandomCounter) printSummary() {
	// log.Info("Counted " ,r.statistic, "add rounds" )
	out, err := json.MarshalIndent(r.statistic, "", "  ")
	if err != nil {
		log.Info(err)
		return
	}
	fmt.Printf(string(out))
}

func NewRandomCounter(counter CounterSystem) *RandomCounter {
	r := new(RandomCounter)
	r.statistic = RandomCounterStatistic{}
	r.setCounter(counter)
	return r
}
