package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// In the 20×20 grid below, four numbers along a diagonal line have been marked in red.
//
// [Puzzle]
//
// The product of these numbers is 26 × 63 × 78 × 14 = 1788696.
//
// What is the greatest product of four adjacent numbers in the same direction (up, down, left, right, or diagonally) in the 20×20 grid?

func main() {
	data, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	var numberGrid [][]int
	fileContent := strings.Split(string(data), "\n")
	for _, fileRow := range fileContent {
		var gridRow []int
		numberStr := strings.Split(fileRow, " ")
		for _, intString := range numberStr {
			number, _ := strconv.Atoi(intString)
			gridRow = append(gridRow, number)
		}
		numberGrid = append(numberGrid, gridRow)
	}

	var maxSum int
	var a, b, c, d int
	var direction string

	// horizontal
	for row := range numberGrid {
		for x := 0; x < 16; x++ {
			sum := numberGrid[row][x] * numberGrid[row][x+1] * numberGrid[row][x+2] * numberGrid[row][x+3]

			if sum > maxSum {
				maxSum = sum
				a, b, c, d = numberGrid[row][x], numberGrid[row][x+1], numberGrid[row][x+2], numberGrid[row][x+3]
				direction = "horizontal"
			}
		}
	}

	// Vertical
	for col := range numberGrid[0] {
		for row := 0; row < 16; row++ {
			sum := numberGrid[row][col] * numberGrid[row+1][col] * numberGrid[row+2][col] * numberGrid[row+3][col]
			if sum > maxSum {
				maxSum = sum
				a, b, c, d = numberGrid[row][col], numberGrid[row+1][col], numberGrid[row+2][col], numberGrid[row+3][col]
				direction = "vertical"
			}

		}

	}

	// diagonal: left*up to right*down
	for r := 0; r <= 16; r++ {
		for c := 0; c <= 16; c++ {
			sum := numberGrid[r][c] * numberGrid[r+1][c+1] * numberGrid[r+2][c+2] * numberGrid[r+3][c+3]
			if sum > maxSum {
				maxSum = sum
				a, b, c, d = numberGrid[r][c], numberGrid[r+1][c+1], numberGrid[r+2][c+2], numberGrid[r+3][c+3]
				direction = "diagonal1"
			}
		}
	}

	// diagnoal2: right*up to left*down
	for row := 3; row <= 19; row++ {
		for col := 0; col <= 16; col++ {
			sum := numberGrid[row][col] * numberGrid[row-1][col+1] * numberGrid[row-2][col+2] * numberGrid[row-3][col+3]
			if sum > maxSum {
				maxSum = sum
				a, b, c, d = numberGrid[row][col], numberGrid[row-1][col+1], numberGrid[row-2][col+2], numberGrid[row-3][col+3]
				direction = "diagonal2"
			}
		}
	}

	fmt.Println(a, b, c, d)
	fmt.Println("Solution found in", direction)
	fmt.Println("011/ Solution:", maxSum)
}
