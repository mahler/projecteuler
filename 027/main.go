package main

import (
	"fmt"
	"math"
)

func main() {
	var maxCounter, aInMax, bInMax int
	maxCounter = 0
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
		inner:
			for n := 0; ; n++ {
				if !isPrime(n*n + a*n + b) {
					if n > maxCounter {
						maxCounter, aInMax, bInMax = n, a, b
					}
					break inner
				}
			}
		}
	}

	fmt.Println("027/ Find the product of the coefficients,  and , for the quadratic expression ")
	fmt.Println("that produces the maximum number of primes for consecutive values of , starting with .")
	fmt.Println(aInMax * bInMax)
}

// https://www.thepolyglotdeveloper.com/2016/12/determine-number-prime-using-golang/
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
