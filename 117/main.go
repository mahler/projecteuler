package main

import (
	"fmt"
)

var table map[int]int64 = make(map[int]int64)

func main() {
	fmt.Println("117/ How many ways can a row measuring fifty units in length be tiled?")
	fmt.Println(tile(50))
}

func tile(n int) int64 {
	if n < 2 {
		return 1
	}

	if answer, ok := table[n]; ok {
		return answer
	}

	answer := int64(1) //empty tiling

	for i := 0; i < n; i++ {
		if n-i >= 4 {
			answer += tile(n - i - 4)
		}
		if n-i >= 3 {
			answer += tile(n - i - 3)
		}
		if n-i >= 2 {
			answer += tile(n - i - 2)
		}
	}

	table[n] = answer
	return answer
}
