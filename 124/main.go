package main

import (
	"fmt"
	"math"
)

const (
	max    = 100000
	target = 10000
)

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

func main() {
	radTable := make(map[int64][]int64)

	for i := int64(2); i <= max; i++ {
		factor := Factor(i)
		factor = RemoveDuplicates(factor)

		rad := int64(1)

		for _, fac := range factor {
			rad *= fac
		}

		radTable[rad] = append(radTable[rad], i)

	}

	current := 1
	var result int64
	for i := int64(0); i <= max; i++ {
		if answer, ok := radTable[i]; ok {

			if current < target && current+len(answer) >= target {
				pos := target - current - 1
				result = answer[pos]
			}
			current += len(answer)
		}
	}
	fmt.Println("124/ If rad(n) is sorted for 1 ≤ n ≤ 100000, find E(10000)")
	fmt.Println(result)

}

func Prime(n int64) int64 {

	if n < 1 {
		return 0
	}

	primeTable[1] = 2
	primeTable[2] = 3

	if n < primeTableLength && primeTable[n] != 0 {
		return primeTable[n]
	}

	i := Prime(n-1) + 1

	for !IsPrime(i) {
		i++
	}

	if i < primeTableLength {
		primepilist[i] = n
	}

	if n < primeTableLength {
		primeTable[n] = i
	}
	return i
}

func Factor(n int64) []int64 {
	var answer = []int64{}

	current := n

	i := int64(1)
	for !IsPrime(current) {
		if current%Prime(i) == 0 {
			answer = append(answer, Prime(i))
			current = current / Prime(i)
			i = 0
		}
		i++
	}

	answer = append(answer, current)

	return answer
}

func IsPrime(n int64) bool {

	if n == 1 {
		return false
	}

	end := int64(math.Sqrt(float64(n)))

	//If we start computing beyond the table, this is stupid
	for i := int64(1); Prime(i) <= end && i < primeTableLength; i++ {
		if n%Prime(i) == 0 {
			return false
		}
	}

	//If we need to pass the end of the prime table brute force
	if end > lastPrime {
		for i := int64(lastPrime); i <= end; i++ {
			if n%i == 0 {
				return false
			}
		}

	}

	return true
}

func RemoveDuplicates(input []int64) []int64 {

	if len(input) == 1 {
		return input
	}

	if input[0] == input[1] {
		return RemoveDuplicates(append(input[1:]))
	}

	return append(input[:1], RemoveDuplicates(input[1:])...)

}
