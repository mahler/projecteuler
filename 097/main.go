package main

import (
	"fmt"
	"math/big"
)

// The first known prime found to exceed one million digits was discovered in 1999,
// and is a Mersenne prime of the form 26972593−1; it contains exactly 2,098,960
// digits. Subsequently other Mersenne primes, of the form 2p−1, have been found
// which contain more digits.
//
// However, in 2004 there was found a massive non-Mersenne prime which contains
// 2,357,207 digits: 28433×2^7830457+1.
//
// Find the last ten digits of this prime number.

func main() {
	var prime big.Int
	prime.Exp(big.NewInt(2), big.NewInt(7830457), nil)
	prime.Mul(big.NewInt(28433), &prime)
	prime.Add(big.NewInt(1), &prime)
	theNumber := fmt.Sprintf("%v", prime.String())

	fmt.Println("097/ Find the last ten digits of this prime number")
	fmt.Println(theNumber[len(theNumber)-10:])
}
