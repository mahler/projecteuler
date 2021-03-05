package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const infinity = 999999

// https://projecteuler.net/problem=82
func main() {

	data := importFile("p082_matrix.txt")

	matrix := make([][]int, len(data))

	for j, line := range data {
		words := strings.Split(line, ",")
		row := make([]int, len(words))
		for i, word := range words {
			row[i], _ = strconv.Atoi(word)
		}
		matrix[j] = row
	}

	sumtrix := make([][]int, len(matrix))
	for i := range sumtrix {
		sumtrix[i] = make([]int, len(matrix[0]))
		sumtrix[i][0] = matrix[i][0]
	}

	for currentColumn := 1; currentColumn < len(matrix[0]); currentColumn++ {

		for i := 0; i < len(matrix); i++ {
			winner := infinity
			for j := 0; j < len(matrix); j++ {
				//now we compute the cost of getting to (i, column) from (j, column -1)
				cost := matrix[i][currentColumn]
				cost += sumtrix[j][currentColumn-1]
				for k := i; sign(j-i)*(j-k) > 0; k += sign(j - i) {
					cost += matrix[k][currentColumn-1]
				}

				if cost < winner {
					winner = cost
				}

			}

			sumtrix[i][currentColumn] = winner

		}
	}

	best := infinity

	for i := 0; i < len(sumtrix); i++ {
		if sumtrix[i][len(sumtrix[i])-1] < best {
			best = sumtrix[i][len(sumtrix[i])-1]
		}
	}

	fmt.Println("082/ the minimal path sum from the left column to the right column in matrix")
	fmt.Println(best)
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}

func importFile(filename string) []string {
	// read whole the file
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var output []string
	currentline := ""

	for _, char := range b {
		if char == 10 {
			output = append(output, currentline)
			currentline = ""
		} else {
			currentline += string(char)
		}
	}

	if currentline != "" {
		output = append(output, currentline)
	}

	return output

}
