package main

import "fmt"

// The prime 41, can be written as the sum of six consecutive primes:
//
// 41 = 2 + 3 + 5 + 7 + 11 + 13
// This is the longest sum of consecutive primes that adds to a prime below one-hundred.
//
// The longest sum of consecutive primes below one-thousand that adds to a prime,
// contains 21 terms, and is equal to 953.
//
// Which prime, below one-million, can be written as the sum of the most consecutive primes?

const belowLimit int = 1000000

func main() {
	arr := make([]bool, belowLimit)
	arr[0], arr[1] = true, true
	count := 2
	prime := 3
	var k int
	finished := false
	for !finished {
		for k = 2 * prime; k < len(arr); k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < len(arr) && arr[k]; k += 2 {
		}
		if k < len(arr) {
			prime = k
			count++
		} else {
			finished = true
		}
	}
	primes := make([]int, count)
	index := 1
	primes[0] = 2
	for i := 0; i < len(arr); i++ {
		if !arr[i] && i%2 != 0 {
			primes[index] = i
			index++
		}
	}

	var answer, maxCount, tmp int
	for i := range primes {
		tmp = primes[i]
		count = 0
	inner2:
		for j := i + 1; j < len(primes); j++ {
			tmp += primes[j]
			count++
			if tmp > len(arr)-1 {
				break inner2
			}
			if tmp%2 != 0 && !arr[tmp] && count > maxCount {
				maxCount = count
				answer = tmp
			}
		}
	}
	fmt.Println("050/ the sum of the most consecutive primes?")
	fmt.Println(answer)
}
