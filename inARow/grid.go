package inARow

import "fmt"

func updateGrid(grid [3][3]Val, gridPosition int, operator string) [3][3]Val {
	position := 0

	for rowNr, row := range grid {
		for cellNr, cell := range row {
			if position == gridPosition && !cell.isUsed {

				grid[rowNr][cellNr].isUsed = true
				grid[rowNr][cellNr].sign = operator
			}
			position++
		}
	}
	return grid
}

func printGrid(grid [3][3]Val) {
	position := 0
	for _, row := range grid {

		for _, cell := range row {

			if cell.isUsed {
				fmt.Print(" ", cell.sign)
			} else {
				fmt.Print(" ", position)
			}

			position++
		}

		fmt.Println("")
	}
}

func getGrid() (grid [3][3]Val) {
	return [3][3]Val{
		{Val{}, Val{}, Val{}},
		{Val{}, Val{}, Val{}},
		{Val{}, Val{}, Val{}},
	}
}
