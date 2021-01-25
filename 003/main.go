package main

import "fmt"

func main() {
	target := 600851475143
	// target := 13195
	var primefactors []int
	// Get the number of 2s that divide n
	for target%2 == 0 {
		primefactors = append(primefactors, 2)
		target = target / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= target; i = i + 2 {
		// while i divides n, append i and divide n
		for target%i == 0 {
			primefactors = append(primefactors, i)
			target = target / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if target > 2 {
		primefactors = append(primefactors, target)
	}

	// fmt.Println(primefactors)
	// Print largest primefactor:
	fmt.Println("003/", primefactors[len(primefactors)-1])
}
