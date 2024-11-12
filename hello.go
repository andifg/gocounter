package main

import (
	//"fmt"
	//"math/rand/v2"
	"playground/hello/counter"
	log "github.com/sirupsen/logrus"
	"os"
)

func init(){
  log.SetFormatter(&log.JSONFormatter{})
  log.SetOutput(os.Stdout)
  log.SetLevel(log.InfoLevel)

}

func main() {
	log.Info("###############")
	counterInstance := counter.NewCounter()
	randomCounter := counter.NewRandomCounter(counterInstance)
	randomCounter.PlayRandomRounds()
}
