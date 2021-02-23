package main

import (
	"fmt"
)

type tree struct {
	parent *tree
	value  int
}

const searchLength = 200 + 1

// The most naive way of computing n15 requires fourteen multiplications:
//
// n × n × ... × n = n15
//
// But using a "binary" method you can compute it in six multiplications:
//
// n × n = n^2
// n^2 × n^2 = n^4
// n^4 × n^4 = n^8
// n^8 × n^4 = n^12
// n^12 × n^2 = n^14
// n^14 × n = n^15
//
// However it is yet possible to compute it in only five multiplications:
//
// n × n = n^2
// n^2 × n = n^3
// n^3 × n^3 = n^6
// n^6 × n^6 = n^12
// n^12 × n^3 = n^15
//
// We shall define m(k) to be the minimum number of multiplications to compute nk; for example m(15) = 5.
//
// For 1 ≤ k ≤ 200, find ∑ m(k).

func main() {
	table := [searchLength]int{}

	haveseen := make(map[int]bool)
	haveseen[1] = true
	table[1] = 0

	current := make([]*tree, 1)
	current[0] = &tree{value: 1}

	for i := 0; i < 14; i++ {

		next := make([]*tree, 0)
		for _, leaf := range current {
			n := leaf.value

			ancestor := leaf
			temp := make([]*tree, 0)
			for ancestor != nil {
				consider := n + ancestor.value
				if !haveseen[consider] {

					if consider < searchLength {
						table[consider] = i + 1

					}
					haveseen[consider] = true
					temp = append([]*tree{{leaf, consider}}, temp...)
				}
				ancestor = ancestor.parent
			}
			next = append(next, temp...)
		}

		current = next

	}

	total := 0
	for i := 1; i < searchLength; i++ {
		total += table[i]
	}

	fmt.Println("122/ For 1 ≤ k ≤ 200, find ∑ m(k)")
	fmt.Println(total - 2)
}
