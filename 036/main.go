package main

import (
	"fmt"
	"strconv"
)

// The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.
//
// Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
//
// (Please note that the palindromic number, in either base, may not include leading zeros.)
func main() {
	dual_palindromes := []int{}
	for i := 0; i < 1000000; i++ {
		if is_dual_palindrome(i) {
			dual_palindromes = append(dual_palindromes, i)
		}
	}

	var sum int
	for _, v := range dual_palindromes {
		sum += v
	}

	fmt.Println("036/ The sum of all numbers, less than one million, which are palindromic in base 10 and base 2?")
	fmt.Println(sum)
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

func is_dual_palindrome(i int) bool {
	// palindrome, base 10:
	b10 := strconv.FormatInt(int64(i), 10)
	// palindrome, base 2
	b02 := strconv.FormatInt(int64(i), 2)
	return is_palindrome(b10) && is_palindrome(b02)
}
