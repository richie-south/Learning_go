package main

import (
	"fmt"
)

import (
	"math"
)

// input out put

func main() {

	/*fmt.Println("true && false = ", true && false)

	  // skapar variable i och asignar den till sin typ
		i := 1

		for i <= 10 {
			fmt.Println(i)
			i++
		}*/

	/*yourAge := 18

	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else if yourAge >= 18 {
		fmt.Println("You can't drive")
	} else {
		fmt.Println("You can't drive")
	}

	switch yourAge {
	case 16:
		fmt.Println("16")
	case 18:
		fmt.Println("18")
	default:
	case 16:
		fmt.Println("default")
	}*/
	/*
		// arrays

		// array of 5 length in float64
		var favNums [2]float64
		favNums[0] = 1.1
		favNums[1] = 1.2

		fmt.Println(favNums[1])

		favNums2 := [2]float64{1.1, 1.2}

		// loop array
		for i, value := range favNums2 {
			fmt.Println(value, i)

		}*/

	/*// slice = like array bu no pre defined size
	numSlice := []int{4, 5, 6, 7, 1}

	numSlice2 := numSlice[2:4] // slice an array (like javascrtip slice) from index 2 to 4
	// numSlice[:2] // auto start at 0 =to numSlice[0:2]
	// numSlice[2:] // start at 2 and gets all items after

	numSlice3 := make([]int, 5, 10)

	copy(numSlice3, numSlice)*/

	/*const pi float64 = 3.14

		var (
			varA = 1
			varB = 2
		)
	  var myName string = "Richard"

	  fmt.Pintln(len(myName))

	  fmt.Println(myName + " Hej")


	  var isOver40 bool = true*/

	/*fmt.Println("Hello world")

		var age int = 40

		var favNum float64 = 1.6180339

		randNum := 1

	  var numOne = 1.000
	  var num99 = .9999

	  fmt.Println(numOne - num99)*/
	// len = gets length
	/*
		// map, key value pair

		presAge := make(map[string]int)

		presAge["richard"] = 42

		fmt.Println(presAge["richard"])

		delete(presAge, "richard") // removes key*/

	// using functions
	listNums := []float64{1, 2, 3, 4, 5}
	fmt.Println("sum : ", addThemUp(listNums))

	//num1, num2 := next2Values(5)

	fmt.Println(subtractThem(1, 2, 3, 4, 5))

	// clojure
	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		return num3
	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

	// recursion
	fmt.Println(factorial(3))

	// defer will rund at absolutly end of function
	defer printTwo()
	printOne()

	fmt.Println(saveDiv(10, 0))

	demPanic()

	x := 0
	// cahanges value of x not x variable (new x value not changis x variable)
	changeXVal(x)

	// cahnges x memory value (changes x variable value)
	changeXValNow(&x)

	// struct
	rect1 := Rectangle{
		leftX:  0,
		topY:   50,
		height: 10,
		width:  10,
	}

	fmt.Println("rec is", rect1.width, "wide")
	// rect1 := Rectangle{ 0, 50, 10, 10} // also works
	fmt.Println("area of rec =", rect1.area())

	// using interface
	rect := Rectangle2{20, 50}
	circ := Circle{4}

	fmt.Println("rect Area =", getArea(rect))
	fmt.Println("circ Area =", getArea(circ))
}

// interface
type Shape interface {
	area() float64
}

type Rectangle2 struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle2) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

// structs

type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

// ads function to structs (typ :P )
func (rec *Rectangle) area() float64 {
	return rec.width * rec.height
}

// some funcs

// changes memory value of x
func changeXValNow(x *int) {
	*x = 2
}

// pointer
func changeXVal(x int) {
	x = 2
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("PANIC") // sends message to recover
}

func saveDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover()) // swalows errors, continue exec even if error
	}() // auto run

	solution := num1 / num2
	return solution
}

func printOne() {
	fmt.Println(1)
}

func printTwo() {
	fmt.Println(2)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)

}

func subtractThem(args ...int) int {
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}
	return finalValue
}

// return 2 values from function
func next2Values(number int) (int, int) {
	return number + 1, number + 2
}

func addThemUp(numbers []float64) float64 {
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
