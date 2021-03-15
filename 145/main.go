package main

import (
	"fmt"
	"strconv"
)

// Some positive integers n have the property that the sum [ n + reverse(n) ] consists entirely of odd
// (decimal) digits. For instance, 36 + 63 = 99 and 409 + 904 = 1313. We will call such numbers reversible;
// so 36, 63, 409, and 904 are reversible. Leading zeroes are not allowed in either n or reverse(n).
//
// There are 120 reversible numbers below one-thousand.
//
// How many reversible numbers are there below one-billion (109)?
func main() {
	total := 0
	for i := int64(0); i < 100000000; i++ {
		if isOdd(i+IntReverse(i)) && i%10 != 0 {
			total++
		}

	}

	fmt.Println("145/ How many reversible numbers are there below one-billion (10^9)?")
	fmt.Println(total)
}

func isOdd(n int64) bool {

	a := strconv.FormatInt(n, 10)

	for i := 0; i < len(a); i++ {
		x, _ := strconv.Atoi(a[i : i+1])
		if x%2 == 0 {
			return false
		}
	}
	return true
}

func IntReverse(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	s = StringReverse(s)
	x, _ := strconv.ParseInt(s, 10, 64)
	return x
}

func StringReverse(a string) string {
	b := ""
	for i := len(a) - 1; i >= 0; i-- {
		b += string(a[i])
	}
	return b
}
