package main

import (
	"fmt"
)

var memo map[[2]int]int64

// Working from left-to-right if no digit is exceeded by the digit to its left it is called an increasing number; for example, 134468.
//
// Similarly if no digit is exceeded by the digit to its right it is called a decreasing number; for example, 66420.
//
// We shall call a positive integer that is neither increasing nor decreasing a "bouncy" number; for example, 155349.
//
// As n increases, the proportion of bouncy numbers below n increases such that there are only 12951 numbers below
// ne-million that are not bouncy and only 277032 non-bouncy numbers below 1010.
//
// How many numbers below a googol (10100) are not bouncy?
func nonbouncy(length int) int64 {
	total := int64(0)
	for i := 1; i <= length; i++ {
		total += increasing(i, 1) //the increasing numbers
		for j := 1; j <= i; j++ {
			total += increasing(j, 1) //the decreasing numbers (potential zeroes)
		}

		total = total - 9 //the constant numbers
	}
	return total
}

func main() {
	memo = make(map[[2]int]int64)
	fmt.Println("113/ How many numbers below a googol (10^100) are not bouncy?")
	fmt.Println(nonbouncy(100))

}

func increasing(length, start int) int64 {
	if length == 1 {
		return 10 - int64(start)
	}

	if answer, ok := memo[[2]int{length, start}]; ok {
		return answer
	}

	answer := int64(0)
	for i := start; i < 10; i++ {
		answer += increasing(length-1, i)
	}

	memo[[2]int{length, start}] = answer

	return answer

}
