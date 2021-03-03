package main

import (
	"fmt"
)

// It turns out that 12 cm is the smallest length of wire that can be bent to form an integer sided right angle triangle in exactly one way, but there are many more examples.
//
// 12 cm: (3,4,5)
// 24 cm: (6,8,10)
// 30 cm: (5,12,13)
// 36 cm: (9,12,15)
// 40 cm: (8,15,17)
// 48 cm: (12,16,20)
//
// In contrast, some lengths of wire, like 20 cm, cannot be bent to form an integer sided right angle triangle, and other lengths allow more than one solution to be found; for example, using 120 cm it is possible to form exactly three different integer sided right angle triangles.
//
// 120 cm: (30,40,50), (20,48,52), (24,45,51)
//
// Given that L is the length of the wire, for how many values of L ≤ 1,500,000 can exactly one integer sided right angle triangle be formed?
func main() {
	max := 1500000
	set := make(map[int]int)
	for m := 1; (2*m*m)+(2*m) < max; m++ {

		for n := m%2 + 1; n < m; n += 2 {
			if greatestCommonDivisor(int64(m), int64(n)) == 1 && (m-n)%2 == 1 {
				for k := 1; k*((2*m*m)+(2*m*n)) < max; k++ {
					set[k*((2*m*m)+(2*m*n))]++
				}
			}
		}
	}

	total := 0
	for _, sum := range set {
		if sum == 1 {
			total++
		}
	}
	fmt.Println("075/ Given that L is the length of the wire, for how many values of L ≤ 1,500,000 can exactly one integer sided right angle triangle be formed?")
	fmt.Println(total)
}

func greatestCommonDivisor(x, y int64) int64 {
	for x != 0 && y != 0 {
		if x > y {
			x %= y
		} else {
			y %= x
		}
	}

	if x > y {
		return x
	} else {
		return y
	}
}
