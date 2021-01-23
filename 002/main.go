package main

import "fmt"

// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
//
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

func main() {

	var numbers []int

	numbers = append(numbers, 1)
	numbers = append(numbers, 2)

	for {
		n := len(numbers)
		nextNumber := numbers[n-2] + numbers[n-1]

		if nextNumber > 4000000 {
			break
		} else {
			numbers = append(numbers, nextNumber)
		}

	}

	sum := 0

	for n := range numbers {
		if numbers[n]%2 == 0 {
			sum += numbers[n]
		}
	}

	fmt.Println("Result:", sum)

}
