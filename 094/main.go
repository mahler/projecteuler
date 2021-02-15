package main

import (
	"fmt"
)

// It is easily proved that no equilateral triangle exists with integral length sides and integral area. However,
// the almost equilateral triangle 5-5-6 has an area of 12 square units.
//
// We shall define an almost equilateral triangle to be a triangle for which two sides are equal and the third
// differs by no more than one unit.
//
// Find the sum of the perimeters of all almost equilateral triangles with integral side lengths and area and whose
// perimeters do not exceed one billion (1,000,000,000).

func main() {
	var maxP, max, total int64

	maxP = 1000000000
	max = 50000

	for n := int64(1); n < max; n++ {
		for m := n + 1; m < max; m += 2 {
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n

			if 2*a == c+1 || 2*a == c-1 {
				perimeter := 2 * (a + c)
				if perimeter <= maxP {
					total += perimeter
				}
			}

			if 2*b == c+1 || 2*b == c-1 {
				perimeter := 2 * (b + c)
				if perimeter <= maxP {
					total += perimeter
				}
			}
		}
	}

	fmt.Println("094/ Find the sum of the perimeters of all almost equilateral triangles with integral")
	fmt.Println("side lengths and area and whose perimeters do not exceed one billion")
	fmt.Println(total)
}
