package main

import (
	"container/list"
	"fmt"
	"math"
)

// A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example,
// the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
//
// A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.
//
// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24.
// By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper
// limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.
//
// Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

const limit = 28124

func main() {
	abunList := list.New()
	writableInt := make([]bool, limit)
	for i := 1; i < limit; i++ {
		writableInt[i] = false
		if i < sumOfProperDivisors(i) {
			abunList.PushBack(i)
		}
	}

	for i := abunList.Front(); i != nil; i = i.Next() {
		for j := i; j != nil; j = j.Next() {
			if sum := i.Value.(int) + j.Value.(int); sum < limit {
				writableInt[sum] = true
			}
		}
	}

	sum := 0
	for i := 1; i < limit; i++ {
		if !writableInt[i] {
			sum += i
		}
	}
	fmt.Println()
	fmt.Println("023/ the sum of all the positive integers which cannot be written as the sum of two abundant numbers?")
	fmt.Println(sum)
}

func sumOfProperDivisors(input int) int {
	limit, sum := int(math.Sqrt(float64(input))), 0
	for i := 1; i <= limit; i++ {
		if input%i == 0 {
			if i == input/i {
				sum += i
			} else {
				sum += i + (input / i)
			}
		}
	}
	return sum - input
}
