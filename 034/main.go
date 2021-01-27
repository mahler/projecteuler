package main

import "fmt"

// 145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
//
// Find the sum of all numbers which are equal to the sum of the factorial of their digits.
//
// Note: As 1! = 1 and 2! = 2 are not sums they are not included.

func main() {
	factorials := make([]int, 10)
	factorials[0] = 1
	for i := 1; i <= 9; i++ {
		factorials[i] = factorials[i-1] * i
	}

	pow10 := 10
	limit := factorials[9]
	for limit > pow10 {
		limit *= 2
		pow10 *= 10
	}

	sum := 0
	var localSum, j int
	for i := 10; i < limit; i++ {
		localSum, j = 0, i
		for j > 9 {
			localSum += factorials[j%10]
			j /= 10
		}
		localSum += factorials[j]

		if i == localSum {
			sum += i
		}
	}
	fmt.Println("034/ the sum of all numbers which are equal to")
	fmt.Println("the sum of the factorial of their digits")
	fmt.Println(sum)
}
