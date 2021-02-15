package main

import (
	"fmt"
	"strconv"
)

const tablesize = 10000000

var table [tablesize]int

// A number chain is created by continuously adding the square of the digits in a number to form a new number until it has been seen before.
//
// For example,
//
// 44 → 32 → 13 → 10 → 1 → 1
// 85 → 89 → 145 → 42 → 20 → 4 → 16 → 37 → 58 → 89
//
// Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop. What is most amazing is that EVERY starting number will eventually arrive at 1 or 89.
//
// How many starting numbers below ten million will arrive at 89?
func main() {

	table[1] = 1
	table[89] = 89

	var total int

	for i := int64(1); i <= 10000000; i++ {
		if goes(i) == 89 {
			total++
		}
	}

	fmt.Println("092/ How many starting numbers below ten million will arrive at 89?")
	fmt.Println(total)
}

func next(n int64) (sum int64) {
	sum = 0
	word := strconv.FormatInt(n, 10)
	for i := 0; i < len(word); i++ {
		num, _ := strconv.Atoi(string(word[i]))
		numb := int64(num)
		sum += numb * numb
	}
	return
}

func goes(n int64) int {
	if n < tablesize-1 && table[n] != 0 {
		return table[n]
	}

	answer := goes(next(n))

	if n < tablesize-1 {
		table[n] = answer
	}

	return answer
}
