package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.
//
//    3
//   7 4
//  2 4 6
// 8 5 9 3
//
// That is, 3 + 7 + 4 + 9 = 23.
//
// Find the maximum total from top to bottom in triangle.txt, a 15K text file containing a triangle with one-hundred rows.

var arr [][]int
var bests [][]int

func main() {
	fileBuf, _ := os.ReadFile("p067_triangle.txt")
	fileStr := strings.Trim(string(fileBuf), "")
	oneDArrStr := strings.Split(fileStr, "\n")
	var line []string
	arr = make([][]int, 100)
	bests = make([][]int, len(arr))
	for i := range arr {
		line = strings.Split(oneDArrStr[i], " ")
		arr[i] = make([]int, len(line))
		bests[i] = make([]int, len(arr[i]))
		for j := range line {
			arr[i][j], _ = strconv.Atoi(line[j])
			bests[i][j] = -1
		}
	}

	fmt.Println("067/ The maximum total from top to bottom:")
	fmt.Println(bestSum(0, 0))
}

func bestSum(i, j int) int {
	var result int
	if bests[i][j] != -1 {
		result = bests[i][j]
	} else if i == len(arr)-1 {
		result = arr[i][j]
		return arr[i][j]
	} else {
		sumLeft, sumRight := bestSum(i+1, j), bestSum(i+1, j+1)
		if sumLeft > sumRight {
			result = arr[i][j] + sumLeft
		} else {
			result = arr[i][j] + sumRight
		}
	}
	bests[i][j] = result
	return result
}
