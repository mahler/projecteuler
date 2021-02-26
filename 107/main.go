package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const log2of40 = 6

func main() {
	data := importFile("network.txt")
	n := len(data)
	test := make([][]int, n)

	for i, line := range data {
		test[i] = make([]int, n)
		words := strings.Split(line, ",")
		for j, word := range words {
			number, _ := strconv.Atoi(word)
			test[i][j] = number
		}
	}

	max := 9999999
	min := 0

	unoptimized := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			unoptimized += test[i][j]
		}
	}

	for i := 0; i < n*n; i++ {

		min = 0
		var testi, testj int

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if test[i][j] < max && test[i][j] > min {
					min = test[i][j]
					testi, testj = i, j
				}
			}
		}

		test[testi][testj] = 0
		test[testj][testi] = 0

		if !isConnected(test) {
			max = min
			test[testi][testj] = min
			test[testj][testi] = min

		}

	}

	total := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			total += test[i][j]
		}
	}

	fmt.Println("107/ find the maximum saving which can be achieved by removing")
	fmt.Println("redundant edges whilst ensuring that the network remains connected")
	fmt.Println(unoptimized - total)

}

func isConnected(A [][]int) bool {
	n := len(A)
	clone := make([][]int, n)
	for i := 0; i < n; i++ {
		clone[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				clone[i][j] = 1
			} else {
				if A[i][j] != 0 {
					clone[i][j] = 1
				}
			}
		}
	}

	for i := 0; i < log2of40; i++ {
		clone = sqrIntMatrix(clone)
		clone = normalize(clone)
	}

	for i := 0; i < n; i++ {
		if clone[0][i] == 0 {
			return false
		}
	}
	return true

}

func normalize(A [][]int) [][]int {
	n := len(A)
	clone := make([][]int, n)
	for i := 0; i < n; i++ {
		clone[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if A[i][j] != 0 {
				clone[i][j] = 1
			}
		}
	}
	return clone
}

func importFile(filename string) []string {
	// read whole the file
	b, err := ioutil.ReadFile(filename)
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

func sqrIntMatrix(A [][]int) [][]int {
	n := len(A)
	square := make([][]int, n)
	for i := 0; i < n; i++ {
		square[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			answer := 0
			for k := 0; k < n; k++ {
				answer += A[j][k] * A[k][i]
			}
			square[j][i] = answer
		}
	}

	return square
}
