package main

import (
	"fmt"
	"math/big"
)

const top = 1000

func main() {
	frac := big.NewRat(2, 1)
	n, d := frac.Num(), frac.Denom()
	table := big.NewInt(0)

	two := big.NewRat(2, 1)

	var total int

	for i := 0; i < top; i++ {
		frac.Add(two, frac.Inv(frac))

		if len(table.Add(n, d).String()) > len(n.String()) {
			total++
		}
	}

	fmt.Println("057/ In the first one-thousand expansions, how many fractions")
	fmt.Println("contain a numerator with more digits than the denominator?")
	fmt.Printf("%d\n", total)

}
