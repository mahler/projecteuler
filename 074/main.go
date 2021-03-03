package main

import (
	"fmt"
	"strconv"
)

const tablesize = 10000000

var table [tablesize]int

// The number 145 is well known for the property that the sum of the factorial of its digits is equal to 145:
//
// 1! + 4! + 5! = 1 + 24 + 120 = 145
//
// Perhaps less well known is 169, in that it produces the longest chain of numbers that link back to 169; it turns out that there are only three such loops that exist:
//
// 169 → 363601 → 1454 → 169
// 871 → 45361 → 871
// 872 → 45362 → 872
//
// It is not difficult to prove that EVERY starting number will eventually get stuck in a loop. For example,
//
// 69 → 363600 → 1454 → 169 → 363601 (→ 1454)
// 78 → 45360 → 871 → 45361 (→ 871)
// 540 → 145 (→ 145)
//
// Starting with 69 produces a chain of five non-repeating terms, but the longest non-repeating chain with a starting number below one million is sixty terms.
//
// How many chains, with a starting number below one million, contain exactly sixty non-repeating terms?
func main() {
	table[145] = 1
	table[1] = 1
	table[2] = 1
	table[40585] = 1

	table[169] = 3
	table[363601] = 3
	table[1454] = 3

	table[871] = 2
	table[45361] = 2

	table[872] = 2
	table[45362] = 2

	count := 0

	for i := 3; i <= 1000000; i++ {

		if hite(i) == 60 {
			count++
		}

	}
	fmt.Println("074/ How many chains, with a starting number below one million, contain exactly sixty non-repeating terms?")
	fmt.Println(count)
}

func hite(height int) int {
	if height < tablesize && table[height] != 0 {
		return table[height]
	}

	next := 0
	word := strconv.Itoa(height)

	for i := 0; i < len(word); i++ {
		is, _ := strconv.Atoi(word[i : i+1])
		next += factorial(is)
	}

	//fmt.Println(word, "->", next)

	if height < tablesize {
		table[height] = hite(next) + 1
	}

	return hite(next) + 1
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
