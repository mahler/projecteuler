package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Take the number 192 and multiply it by each of 1, 2, and 3:
//
// 192 × 1 = 192
// 192 × 2 = 384
// 192 × 3 = 576
// By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)
//
// The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).
//
//  What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?

func main() {
	/* 9 * (1, 2, 3, 4, 5) is the biggest run possible */
	runs := [4][]int{
		[]int{1, 2},
		[]int{1, 2, 3},
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3, 4, 5},
	}

	// Buffer to capture all found pdProducts
	pandigitalProducts := []string{}

	for i := 1; i < 50000; i++ {
		for _, r := range runs {
			prods := []string{}
			for _, v := range r {
				prods = append(prods, strconv.Itoa(i*v))
			}
			prod := strings.Join(prods, "")
			if is_pandigital(prod) {
				pandigitalProducts = append(pandigitalProducts, prod)
			}
		}
	}

	max := 0
	for _, prod := range pandigitalProducts {
		num, _ := strconv.Atoi(prod)
		if num > max {
			max = num
		}
	}

	fmt.Println("039/ What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?")
	fmt.Printf("Max: %d\n", max)
}

func is_pandigital(s string) bool {
	if len(s) != 9 {
		return false
	}
	set := map[rune]int{}
	for _, r := range s {
		if r == '0' || set[r] > 0 {
			return false
		}
		set[r] = 1
	}
	return true
}
