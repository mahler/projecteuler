package main

import "fmt"

// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//
// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

func main() {
	var a, b, c int

	for a = 1; a < 500; a++ {
		for b = a + 1; b < 1000; b++ {
			c = 1000 - a - b
			if b > c {
				break
			}
			if a*a+b*b == c*c {
				goto found
			}
		}
	}
found:

	fmt.Println("Found numbers:", a, b, c)
	fmt.Println("009/ Product:", a*b*c)
}
