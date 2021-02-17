package main

import (
	"fmt"
)

const top int64 = 1000000000

// Consider the set S(r) of points (x,y) with integer coordinates satisfying |x| + |y| ≤ r.
// Let O be the point (0,0) and C the point (r/4,r/4).
//
// Let N(r) be the number of points B in S(r), so that the triangle OBC has an obtuse angle,
// i.e. the largest angle α satisfies 90°<α<180°.
//
// So, for example, N(4)=24 and N(8)=100.
// What is N(1,000,000,000)?

func main() {
	omega := top / 4
	ans := omega * omega * 24

	aLength := omega + 1
	aOdd := aLength%2 == 1
	a := func(i, x int64) bool { return x*x < i*omega-i*i }
	ans += accum((aLength+1)/2-1, aOdd, a)

	bLength := omega
	bOdd := bLength%2 == 1
	b := func(i, x int64) bool { return 2*x*x-2*x+1 < -2*i-2*i*i+omega+2*i*omega }
	ans += accum((bLength+1)/2-1, bOdd, b)

	fmt.Println("210/ What is N(1,000,000,000)?")
	fmt.Println(ans)
}

func accum(to int64, isOdd bool, eval func(int64, int64) bool) (store int64) {
	point := int64(0)

	for i := int64(0); i <= to; i++ {
		for eval(i, point+1) {
			point++
		}
		store += point
	}
	store *= 4

	if isOdd {
		store -= 2 * point
	}

	return store
}
