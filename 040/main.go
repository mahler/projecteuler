package main

import (
	"fmt"
)

// An irrational decimal fraction is created by concatenating the positive integers:
//
// 0.123456789101112131415161718192021...
//
// It can be seen that the 12th digit of the fractional part is 1.
//
// If dn represents the nth digit of the fractional part, find the value of the following expression.
//
// See https://projecteuler.net/problem=40 for formula

func main() {
	var current, pos, i int
	positions := [7]int{1, 10, 100, 1000, 10000, 100000, 1000000}
	var digits [7]int
	var s string
	for {
		s = fmt.Sprint(i)
		if pos <= positions[current] && positions[current] < pos+len(s) {
			digits[current] = int(s[positions[current]-pos]) - 48
			current++
		}
		pos += len(s)
		if current == 7 {
			break
		}
		i++
	}
	// 	fmt.Println(digits)
	mul := 1
	for _, d := range digits {
		mul *= d
	}

	fmt.Println("040/ the value of the expression")
	fmt.Printf("Product of digits: %d\n", mul)
}
