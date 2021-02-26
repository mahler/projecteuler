package main

import (
	"fmt"
)

const size = 12

// Let S(A) represent the sum of elements in set A of size n. We shall call it a special sum set if for any two non-empty disjoint subsets, B and C, the following properties are true:
//
// 1. S(B) ≠ S(C); that is, sums of subsets cannot be equal.
// 2. If B contains more elements than C then S(B) > S(C).
//
// For this problem we shall assume that a given set contains n strictly increasing elements and it already satisfies the second rule.
//
// Surprisingly, out of the 25 possible subset pairs that can be obtained from a set for which n = 4, only 1 of these pairs need to be tested for equality (first rule). Similarly, when n = 7, only 70 out of the 966 subset pairs need to be tested.
//
// For n = 12, how many of the 261625 subset pairs that can be obtained need to be tested for equality?

func main() {
	count := 0
	compare := make([][][]int, size)

	for i := 1; i < int(IntExp(int64(2), size)); i++ {
		set := binary(i, size)

		blip := make([]int, 0)

		drip := 0
		for a := 0; a < size; a++ {
			if set[a] == 1 {
				drip++
				blip = append(blip, a)
			}
		}

		compare[drip-1] = append(compare[drip-1], blip)

	}

	for i := 1; i < size; i++ {

		for a := 0; a < len(compare[i]); a++ {
			for b := a + 1; b < len(compare[i]); b++ {
				t1 := compare[i][a]
				t2 := compare[i][b]

				elim := false

				for x := 0; x < len(t1); x++ {

					for y := 0; y < len(t1); y++ {

						if t1[x] == t2[y] {
							elim = true
						}

					}
				}

				test := true

				for q := 0; q < len(t1); q++ {
					if t1[q] < t2[q] {
						test = false
						break
					}
				}

				elim = test || elim

				if !elim {
					//fmt.Println(t1, t2)
					count++

				}

			}
		}

	}
	fmt.Println("106/ For n = 12, how many of the 261625 subset pairs that can be obtained need to be tested for equality?")
	fmt.Println(count)
}

func binary(index, size int) []int {
	if size == 0 {
		return []int{}
	}
	return append(binary(index/2, size-1), index%2)

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
