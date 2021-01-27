package main

import (
	"fmt"
	"sort"
)

func main() {
	primenumberList := sieveOfErastothenes(999999999)
	pandigitalMax := 0
	for i := 0; i < len(primenumberList); i++ {
		if pandigitalCheck(primenumberList[i]) {
			if primenumberList[i] > pandigitalMax {
				pandigitalMax = primenumberList[i]
			}
		}
	}
	fmt.Println("041/ The largest n-digit pandigital prime is ")
	fmt.Println(pandigitalMax)
}

// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func sieveOfErastothenes(maxVal int) []int {
	sieve := make([]bool, maxVal)
	primeList := make([]int, 0)

	for i := 3; i < maxVal; i += 2 {
		if sieve[i] == false {
			primeList = append(primeList, i)
		}

		for j := i * 2; j < maxVal; j += i {
			sieve[j] = true
		}
	}
	return primeList
}

func pandigitalCheck(i int) bool {
	digitBuff := make([]int, 0)

	for i != 0 {
		j := (i % 10)
		i /= 10
		digitBuff = append(digitBuff, j)
	}
	sort.Ints(digitBuff)

	for i = 0; i < len(digitBuff); i++ {
		if digitBuff[i] != (i + 1) {
			return false
		}
	}
	return true
}
