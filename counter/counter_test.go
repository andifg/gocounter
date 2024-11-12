package counter

import (
	"testing"
)

func TestAddMethod(t *testing.T) {
	counter := new(Counter)
	counter.Add(3)

	if counter.sum != 3 {
		t.Fatal("add did not work")
	}

}
