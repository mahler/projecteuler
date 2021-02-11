package main

import (
	"fmt"
	"math/big"
)

// If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.
//
// Not all numbers produce palindromes so quickly. For example,
//
// 349 + 943 = 1292,
// 1292 + 2921 = 4213
// 4213 + 3124 = 7337
//
// That is, 349 took three iterations to arrive at a palindrome.
//
// Although no one has proved it yet, it is thought that some numbers, like 196, never
// produce a palindrome. A number that never forms a palindrome through the reverse and
// add process is called a Lychrel number. Due to the theoretical nature of these
// numbers, and for the purpose of this problem, we shall assume that a number is
// Lychrel until proven otherwise. In addition you are given that for every number
// below ten-thousand, it will either (i) become a palindrome in less than fifty
// iterations, or, (ii) no one, with all the computing power that exists, has managed
// so far to map it to a palindrome. In fact, 10677 is the first number to be shown to
// require over fifty iterations before producing a palindrome:
// 4668731596684224866951378664 (53 iterations, 28-digits).
//
// Surprisingly, there are palindromic numbers that are themselves Lychrel numbers;
// the first example is 4994.
//
// How many Lychrel numbers are there below ten-thousand?

func main() {
	num := 0
	is_lychrel(big.NewInt(4994))
	for x := 1; x < 10000; x++ {
		if !is_lychrel(big.NewInt(int64(x))) {
			num++
		}
	}
	fmt.Println("055/ How many Lychrel numbers are there below ten-thousand?")
	fmt.Println(num)
}

func is_lychrel(num *big.Int) bool {
	for i := 0; i <= 50; i++ {
		num.Add(num, reverse(num))
		if is_palindrome(num.String()) {
			return true
		}
	}
	return false
}

func is_palindrome(s string) bool {
	var i, j int
	j = len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func reverse(input *big.Int) *big.Int {
	num := big.NewInt(0)
	num.Set(input)
	toReturn := big.NewInt(0)
	ten := big.NewInt(10)
	for num.String() != "0" {
		temp := big.NewInt(0)
		temp.Mod(num, ten)
		toReturn.Mul(ten, toReturn)
		toReturn.Add(toReturn, temp)
		num.Div(num, ten)
	}
	return toReturn
}
