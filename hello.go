package main

import "fmt"
import "playground/hello/counter"
import "math/rand/v2"

func main() {

	//counter.Testcall()
	fmt.Println("###############")
	t := new(counter.Counter)

	//t.Add(2)
	//t.Add(7)
	//t.Subtract(8)

	rounds := rand.IntN(100)
	fmt.Println("Playing ", rounds, "rounds")

	for num := range rounds {
		fmt.Printf("num: %d \n", num)
		randomNumber := rand.IntN(50)
		t.ExecRandom(randomNumber)
		//fmt.Printf("##############/n")
		//fmt.Printf("End number %d", t.sum)
	}
	
	fmt.Println("#########################")

	counter.ExecRandom2(t)

}

func helloWorld() {
	//fmt.Println("Hello world !")
	//fmt.Println(Expoo)
	//	fmt.Println(Test)
	fmt.Println(Teststring)
	fmt.Println("###############")
}
