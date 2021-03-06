package main

import (
	"fmt"
)

// If we are presented with the first k terms of a sequence it is impossible to say with certainty the value of the next term, as there are infinitely many polynomial functions that can model the sequence.
//
// As an example, let us consider the sequence of cube numbers. This is defined by the generating function,
// un = n3: 1, 8, 27, 64, 125, 216, ...
//
// Suppose we were only given the first two terms of this sequence. Working on the principle that "simple is best" we should assume a linear relationship and predict the next term to be 15 (common difference 7). Even if we were presented with the first three terms, by the same principle of simplicity, a quadratic relationship should be assumed.
//
// We shall define OP(k, n) to be the nth term of the optimum polynomial generating function for the first k terms of a sequence. It should be clear that OP(k, n) will accurately generate the terms of the sequence for n ≤ k, and potentially the first incorrect term (FIT) will be OP(k, k+1); in which case we shall call it a bad OP (BOP).
//
// As a basis, if we were only given the first term of sequence, it would be most sensible to assume constancy; that is, for n ≥ 2, OP(1, n) = u1.
//
// Hence we obtain the following OPs for the cubic sequence:
//
// OP(1, n) = 1	1, 1, 1, 1, ...
// OP(2, n) = 7n−6	1, 8, 15, ...
// OP(3, n) = 6n2−11n+6     	1, 8, 27, 58, ...
// OP(4, n) = n3	1, 8, 27, 64, 125, ...
// Clearly no BOPs exist for k ≥ 4.
//
// By considering the sum of FITs generated by the BOPs (indicated in red above), we obtain 1 + 15 + 58 = 74.
//
// Consider the following tenth degree polynomial generating function:
//
// un = 1 − n + n^2 − n^3 + n^4 − n^5 + n^6 − n^7 + n^8 − n^9 + n^10
//
// Find the sum of FITs for the BOPs.
//

func main() {
	pts := [][2]float64{{1, f(1)}}
	answer := float64(0)

	for i := float64(0); i < 10; i++ {
		answer += lagrange(i+2, pts)
		pts = append(pts, [2]float64{i + 2, f(i + 2)})
	}

	fmt.Println("101/ Find the sum of FITs for the BOPs")
	fmt.Println(int64(answer))
}

func basisPoly(x float64, index int, points [][2]float64) float64 {
	answer := float64(1)
	for i := 0; i < index; i++ {
		answer *= (x - points[i][0]) / (points[index][0] - points[i][0])
	}
	for i := index + 1; i < len(points); i++ {
		answer *= (x - points[i][0]) / (points[index][0] - points[i][0])
	}
	return answer
}

func lagrange(x float64, points [][2]float64) float64 {
	answer := float64(0)
	for i := 0; i < len(points); i++ {
		answer += points[i][1] * basisPoly(x, i, points)
	}
	return answer
}

func f(n float64) float64 {
	return 1 - n + (n * n) - (n * n * n) + (n * n * n * n) - (n * n * n * n * n) + (n * n * n * n * n * n) - (n * n * n * n * n * n * n) + (n * n * n * n * n * n * n * n) - (n * n * n * n * n * n * n * n * n) + (n * n * n * n * n * n * n * n * n * n)
}
