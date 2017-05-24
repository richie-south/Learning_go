package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// strings

	sampString := "Hello World"

	fmt.Println(strings.Contains(sampString, "lo"))
	fmt.Println(strings.Index(sampString, "lo"))
	fmt.Println(strings.Count(sampString, "l"))
	fmt.Println(strings.Replace(sampString, "l", "x", 3))

	csvString := "1,2,3,4,5,6"

	fmt.Println(strings.Split(csvString, ","))

	listOfLetters := []string{"c", "a", "b"}

	sort.Strings(listOfLetters) // modifies original

	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")

	fmt.Println(listOfNums)

	// file

	file, err := os.Create("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This is som text")
	file.Close()

	stream, err := ioutil.ReadFile("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)

	fmt.Println(readString)

	// convert
	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"

	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))

	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)

	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat)

}
