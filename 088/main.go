package main

import (
	"fmt"
)

// A natural number, N, that can be written as the sum and product of a given set of at least two natural numbers,
// {a1, a2, ... , ak} is called a product-sum number: N = a1 + a2 + ... + ak = a1 × a2 × ... × ak.
//
// For example, 6 = 1 + 2 + 3 = 1 × 2 × 3.
//
// For a given set of size, k, we shall call the smallest N with this property a minimal product-sum number.
// The minimal product-sum numbers for sets of size, k = 2, 3, 4, 5, and 6 are as follows.
//
// k=2: 4 = 2 × 2 = 2 + 2
// k=3: 6 = 1 × 2 × 3 = 1 + 2 + 3
// k=4: 8 = 1 × 1 × 2 × 4 = 1 + 1 + 2 + 4
// k=5: 8 = 1 × 1 × 2 × 2 × 2 = 1 + 1 + 2 + 2 + 2
// k=6: 12 = 1 × 1 × 1 × 1 × 2 × 6 = 1 + 1 + 1 + 1 + 2 + 6
//
// Hence for 2≤k≤6, the sum of all the minimal product-sum numbers is 4+6+8+12 = 30; note that 8 is only counted once in the sum.
//
// In fact, as the complete set of minimal product-sum numbers for 2≤k≤12 is {4, 6, 8, 12, 15, 16}, the sum is 61.
//
// What is the sum of all the minimal product-sum numbers for 2≤k≤12000?
func main() {
	recordTable := make(map[int]bool)

	fmt.Println("Patience.. may run for a long time.")

	bottom := 2
	top := 12000

	c := make(chan int)

	for length := bottom; length <= top; length++ {

		go search(length, c)
		//fmt.Println("launched", length)

	}

	for length := bottom; length <= top; length++ {

		answer := <-c
		// fmt.Println(length)
		// fmt.Println(answer)

		recordTable[answer] = true
	}

	total := 0
	for i := range recordTable {
		total += i
	}

	fmt.Println("088/ What is the sum of all the minimal product-sum numbers for 2≤k≤12000?")
	fmt.Println(total)

}

func sum(list []int) (total int) {
	for _, item := range list {
		total += item
	}
	return
}

func prod(list []int) (prod int) {
	prod = 1
	for _, item := range list {
		prod *= item
	}
	return
}

const (
	infinity = 99999999
)

func search(length int, c chan int) {

	limit := 3 * length

	test := make([]int, length)

	inner := func() {
		if sum(test) == prod(test) {
			limit = sum(test)
		}
	}

	var level func(int)

	level = func(lvl int) {

		start := 1

		if lvl > 1 {
			start = test[lvl-2]
		}

		leftprod := int64(1)

		var rightprod int64
		if start >= 2 && 1+length-lvl >= 13 {
			rightprod = infinity
		} else {
			rightprod = IntExp(int64(start), int64(1+length-lvl))
		}

		for test[lvl-1] = start; sum(test[0:lvl])+((length-lvl)*test[lvl-1]) < limit && leftprod*rightprod < int64(limit); test[lvl-1]++ {

			if lvl >= length {
				inner()
			} else {
				level(lvl + 1)
			}

			if lvl > 0 {
				leftprod = int64(prod(test[0 : lvl-1]))
			} else {
				leftprod = 1
			}

			if test[lvl-1] >= 2 && 1+length-lvl >= 13 {
				rightprod = infinity
			} else {
				rightprod = IntExp(int64(test[lvl-1]), int64(1+length-lvl))
			}
		}
	}

	level(1)

	c <- limit
}

func IntExp(a int64, b int64) int64 {
	if a == 0 {
		return 0
	}

	if a > 0 {
		return int64(UInt64Exp(uint64(a), uint64(b)))
	}
	if a < 0 && b%2 == 0 {
		return IntExp(-1*a, b)
	}

	return -1 * IntExp(-1*a, b)

}

func UInt64Exp(a, b uint64) uint64 {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	if b%2 == 0 {
		temp := UInt64Exp(a, b/2)
		return temp * temp
	}
	return a * UInt64Exp(a, b-1)
}
