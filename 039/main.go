package main

import (
	"fmt"
	"math"
)

// If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.
//
// {20,48,52}, {24,45,51}, {30,40,50}
//
// For which value of p ≤ 1000, is the number of solutions maximised?

func main() {
	triplets := [][3]int{}
	for i := 1; i < 500; i++ {
		for j := i; j < 500; j++ {
			k := i*i + j*j

			l := int(math.Sqrt(float64(k)))
			if !(l*l == k) {
				continue
			}

			k = int(math.Sqrt(float64(k)))
			if i+j+k > 1000 {
				break
			}
			triplets = append(triplets, [3]int{i, j, k})
		}
	}
	counts := map[int]int{}
	for _, t := range triplets {
		sum := t[0] + t[1] + t[2]
		counts[sum] = counts[sum] + 1
	}

	// Find maximum
	max, pos := 0, 0
	for k, v := range counts {
		if v > max {
			pos = k
			max = v
		}
	}

	fmt.Println("039/ For which value of p ≤ 1000, is the number of solutions maximised?")
	fmt.Printf("%d triplets add up to %d\n", max, pos)
}
