package main

import (
	"fmt"
	"math"
)

const top = 10000

// Let N be a positive integer and let N be split into k equal parts, r = N/k,
// so that N = r + r + ... + r.
// Let P be the product of these parts, P = r × r × ... × r = r^k.
//
// For example, if 11 is split into five equal parts, 11 = 2.2 + 2.2 + 2.2 + 2.2 + 2.2,
// then P = 2.2^5 = 51.53632.
//
// Let M(N) = Pmax for a given value of N.
//
// It turns out that the maximum for N = 11 is found by splitting eleven into four equal parts
// which leads to Pmax = (11/4)^4; that is, M(11) = 14641/256 = 57.19140625, which is a terminating
// decimal.
//
// However, for N = 8 the maximum is achieved by splitting it into three equal parts,
// so M(8) = 512/27, which is a non-terminating decimal.
//
// Let D(N) = N if M(N) is a non-terminating decimal and D(N) = -N if M(N) is a terminating decimal.
//
// For example, ∑ D(N) for 5 ≤ N ≤ 100 is 2438.
//
// Find ∑ D(N) for 5 ≤ N ≤ 10000.
func main() {
	var sum int
	for N := 5; N <= top; N++ {
		term := M(int64(N))

		if term {
			sum -= N
		} else {
			sum += N
		}
	}

	fmt.Println("183/ Find ∑ D(N) for 5 ≤ N ≤ 10000")
	fmt.Println(sum)
}

func M(n int64) (term bool) {
	try := float64(n) / math.E
	try1 := float64(int64(try))
	try2 := try1 + 1

	test := math.Pow(try2/try1, try1)
	test *= try2 / float64(n)

	var den int64
	if test > 1 {
		den = int64(try1)
	} else {
		den = int64(try2)
	}

	for den%2 == 0 {
		den /= 2
	}

	for den%5 == 0 {
		den /= 5
	}

	gcd := greatestCommonDivisor(n, den)
	den /= gcd

	if den == 1 {
		term = true
	}

	return

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
