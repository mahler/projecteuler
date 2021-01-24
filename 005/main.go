package main

import "fmt"

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

func main() {
	min, max := 1, 20

	// Lowest possible number (with large margin)
	testNumber := max

	for {
		testNumber++
		testFail := false
		for x := min; x <= max; x++ {
			if testNumber%x != 0 {
				testFail = true
			}
		}

		if !testFail {
			break
		}

	}
	fmt.Println("Found number:", testNumber)
}
