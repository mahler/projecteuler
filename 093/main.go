package main

import (
	"fmt"
	"strconv"
)

const zero = .001

// By using each of the digits from the set, {1, 2, 3, 4}, exactly once, and making use of the four
// arithmetic operations (+, −, *, /) and brackets/parentheses, it is possible to form different positive integer targets.
//
// For example,
//
// 8 = (4 * (1 + 3)) / 2
// 14 = 4 * (3 + 1 / 2)
// 19 = 4 * (2 + 3) − 1
// 36 = 3 * 4 * (2 + 1)
//
// Note that concatenations of the digits, like 12 + 34, are not allowed.
//
// Using the set, {1, 2, 3, 4}, it is possible to obtain thirty-one different target numbers of which 36 is the maximum,
// and each of the numbers 1 to 28 can be obtained before encountering the first non-expressible number.
//
// Find the set of four distinct digits, a < b < c < d, for which the longest set of consecutive positive integers, 1 to n,
// can be obtained, giving your answer as a string: abcd.
func main() {
	var longestStreak int
	var answer string

	for a := float64(1); a < 10; a++ {

		for b := a + 1; b < 10; b++ {
			for c := b + 1; c < 10; c++ {
				for d := c + 1; d < 10; d++ {

					ordered := []float64{a, b, c, d}

					table := make(map[int]bool)

					for reorder := 0; reorder < 24; reorder++ {

						for op1 := 0; op1 < 4; op1++ {
							for op2 := 0; op2 < 4; op2++ {
								for op3 := 0; op3 < 4; op3++ {
									for pl1 := 0; pl1 < 3; pl1++ {
										for pl2 := 0; pl2 < 2; pl2++ {
											list := make([]float64, len(ordered))
											copy(list, ordered)
											list = PermuteFloats(reorder, list)
											list = apply(list, pl1, op1)
											list = apply(list, pl2, op2)
											list = apply(list, 0, op3)
											if isInt(list[0]) > 0 {
												table[isInt(list[0])] = true
											}

										}
									}
								}
							}
						}

					}

					var best, streakcount int
					streak := false
					for i := 1; i <= 9*9*9*9; i++ {

						if table[i] {
							if streak {
								streakcount++
							}

							streak = true
						} else {
							streak = false
							if streakcount > best {
								best = streakcount
							}
							streakcount = 0
						}
					}

					if best > longestStreak {
						longestStreak = best
						answer = strconv.Itoa(int(a)) + strconv.Itoa(int(b)) + strconv.Itoa(int(c)) + strconv.Itoa(int(d))
					}
				}
			}
		}
	}

	fmt.Println("093/ the set of four distinct digits, a < b < c < d, for which the longest set of consecutive positive integers")
	fmt.Println(answer)

}

func PermuteFloats(n int, list []float64) []float64 {
	if len(list) == 1 {
		return list
	}

	k := n % len(list)

	first := []float64{list[k]}
	next := make([]float64, len(list)-1)

	copy(next, append(list[:k], list[k+1:]...))

	return append(first, PermuteFloats(n/len(list), next)...)
}

func apply(list []float64, place int, operation int) []float64 {
	answer := make([]float64, len(list)-1)
	copy(answer[:place], list[:place])
	answer[place] = operate(list[place], list[place+1], operation)
	copy(answer[place+1:], list[place+2:])
	return answer
}

func operate(x, y float64, which int) float64 {
	if which == 0 {
		return x + y
	} else if which == 1 {
		return x - y
	} else if which == 2 {
		return x * y
	} else if which == 3 {
		return x / y
	}
	return 0
}

func isInt(a float64) int {
	if a-float64(int(a)) > zero {
		return -1
	}
	return int(a)

}
