package main

import (
	"fmt"
	"strconv"
)

// Find the unique positive integer whose square has the form 1_2_3_4_5_6_7_8_9_0,
// where each “_” is a single digit.
func main() {
	var solution int64

	for i := int64(1000000000); !passTest(i * i); i += 10 {
		solution = i + 10
	}
	fmt.Println("206/ Find the unique positive integer whose square has the form 1_2_3_4_5_6_7_8_9_0, ")
	fmt.Println("where each “_” is a single digit.")
	fmt.Println(solution)
}

func passTest(n int64) bool {
	// Convert to string for easier testing
	w := strconv.FormatInt(n, 10)

	if len(w) != 19 {
		return false
	}

	for i := 1; i < 10; i++ {
		if w[2*i-2:2*i-1] != strconv.Itoa(i) {
			return false
		}
	}

	if w[18:] != "0" {
		return false
	}

	return true
}
