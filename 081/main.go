package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("p081_matrix.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileContent := strings.Split(string(data), "\n")

	var grid [80][80]int
	for i := 0; i < 80; i++ {
		numbers := strings.Split(fileContent[i], ",")
		for j := 0; j < 80; j++ {
			text := numbers[j]
			grid[i][j], _ = strconv.Atoi(text)

		}
	}

	//fmt.Println(grid)

	for i := 78; i > -1; i-- {
		grid[79][i] += grid[79][i+1]
		grid[i][79] += grid[i+1][79]
	}

	for i := 78; i > -1; i-- {
		for j := 78; j > -1; j-- {
			if grid[i+1][j] < grid[i][j+1] {
				grid[i][j] += grid[i+1][j]
			} else {
				grid[i][j] += grid[i][j+1]
			}
		}
	}

	fmt.Println("081/ Find the minimal path sum from the top left to the bottom right by only moving right and down")
	fmt.Println(grid[0][0])
}
