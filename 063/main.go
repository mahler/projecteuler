package main

import (
	"fmt"
	"math/big"
)

// The 5-digit number, 16807=75, is also a fifth power. Similarly, the 9-digit number, 134217728=8^9, is a ninth power.
//
// How many n-digit positive integers exist which are also an nth power?

func main() {
	total := 0
	for i := int64(1); i < 10; i++ {
		for j := int64(1); j < 26; j++ {
			number := big.NewInt(i)
			power := big.NewInt(j)

			temp := new(big.Int).Exp(number, power, nil)
			lenght := len(temp.String())
			l := int64(lenght)

			if l == j {
				total++
			}
		}
	}

	fmt.Println("063/ How many n-digit positive integers exist which are also an nth power?")
	fmt.Println("Total:", total)
}
