package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNumber = 0
var pizzaName = ""

func main() {

	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond * 5000)
	}

}

func makeDough(stringChan chan string) {
	pizzaNumber++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNumber)

	fmt.Println("make dough ad send for sauce")

	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	pizza := <-stringChan
	fmt.Println("Add sause send", pizza, "for toppings")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChan chan string) {

	pizza := <-stringChan

	fmt.Println("Add toppings", pizza, "and ship")
	time.Sleep(time.Millisecond * 10)
}
