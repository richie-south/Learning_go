package inARow

import "fmt"

const (
	PlayerOne = "x"
	PlayerTwo = "o"
)

type Val struct {
	isUsed bool
	sign   string
}

func Start() {
	grid := getGrid()
	isPlayerOne := true
	var player string

	for {
		printGrid(grid)
		input := readInput()

		if isPlayerOne {
			player = PlayerOne
			isPlayerOne = false
		} else {
			player = PlayerTwo
			isPlayerOne = true
		}

		grid = updateGrid(grid, input, player)
	}
}

func readInput() int {
	var i int
	fmt.Scan(&i)
	return i
}
