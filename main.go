package main

import (
	"fmt"

	"github.com/rs222kn/Learning_go/firstClass"
)

func main() {

	fmt.Println("go")
	myArray := []string{"a", "c", "b", "a"}

	resultFilter := firstClass.MyFilter(myArray, func(value string) bool {

		if value == "a" {
			return true
		}
		return false

	})

	fmt.Println(resultFilter)

	resultMap := firstClass.MyMap(myArray, func(value string) string {
		return "#" + value
	})

	fmt.Println(resultMap)
	fmt.Println("Hello world!")
}
