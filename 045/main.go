package main

import (
	"fmt"
)

// https://projecteuler.net/problem=45
func main() {
	var penmap = map[int64]bool{}
	var hexmap = map[int64]bool{}

	for i := int64(2); ; i++ {
		// Calc Triangle number:
		t := (i * (i + 1)) / 2
		// Calc Pentagonal number:
		p := (i * ((3 * i) - 1)) / 2
		// Calc Hexagonal number:
		h := (i * ((2 * i) - 1))

		penmap[p] = true
		hexmap[h] = true

		if penmap[t] && hexmap[t] && t > 40755 {
			fmt.Println("045/ Find the next triangle number that is also pentagonal and hexagonal?")
			fmt.Println(t)
			return
		}

	}
}
