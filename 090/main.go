package main

import (
	"fmt"
)

// From https://projecteuler.net/problem=90:
//
// Each of the six faces on a cube has a different digit (0 to 9) written on it; the same is done to a second cube. By placing the two cubes side-by-side in different positions we can form a variety of 2-digit numbers.
//
// For example, the square number 64 could be formed:
//
//    Cube1: [6] - Cube2: [4] => [6][4]
//
// In fact, by carefully choosing the digits on both cubes it is possible to display all of the square numbers below one-hundred: 01, 04, 09, 16, 25, 36, 49, 64, and 81.
//
// For example, one way this can be achieved is by placing {0, 5, 6, 7, 8, 9} on one cube and {1, 2, 3, 4, 8, 9} on the other cube.
//
// However, for this problem we shall allow the 6 or 9 to be turned upside-down so that an arrangement like {0, 5, 6, 7, 8, 9} and {1, 2, 3, 4, 6, 7} allows for all nine square numbers to be displayed; otherwise it would be impossible to obtain 09.
//
// In determining a distinct arrangement we are interested in the digits on each cube, not the order.
//
// {1, 2, 3, 4, 5, 6} is equivalent to {3, 6, 4, 1, 2, 5}
// {1, 2, 3, 4, 5, 6} is distinct from {1, 2, 3, 4, 5, 9}
//
// But because we are allowing 6 and 9 to be reversed, the two distinct sets in the last example both represent the extended set {1, 2, 3, 4, 5, 6, 9} for the purpose of forming 2-digit numbers.

func main() {
	die1 := trues()
	die2 := trues()
	total := 0

	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				for l := k + 1; l < 10; l++ {

					die1 = trues()
					die1[i] = false
					die1[j] = false
					die1[k] = false
					die1[l] = false
					if die1[6] || die1[9] {
						die1[9] = true
						die1[6] = true
					}

					for I := 0; I < 10; I++ {
						for J := I + 1; J < 10; J++ {
							for K := J + 1; K < 10; K++ {
								for L := K + 1; L < 10; L++ {

									die2 = trues()
									die2[I] = false
									die2[J] = false
									die2[K] = false
									die2[L] = false
									if die2[6] || die2[9] {
										die2[9] = true
										die2[6] = true
									}

									if works(die1, die2) {
										total++

									}

								}
							}
						}
					}
				}

			}
		}
	}

	fmt.Println("090/ How many distinct arrangements of the two cubes allow for all of the square numbers to be displayed?")
	fmt.Println(total / 2)
}

func works(a, b map[int]bool) bool {
	for i := 1; i <= 9; i++ {
		d1 := (i * i) / 10
		d2 := (i * i) % 10

		if (a[d1] && b[d2]) || (a[d2] && b[d1]) {
			//too lazy to negate this
		} else {
			return false
		}
	}
	return true
}

func trues() map[int]bool {
	die := make(map[int]bool)
	for i := 0; i < 10; i++ {
		die[i] = true
	}

	return die
}
