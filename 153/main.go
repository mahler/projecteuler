package main

import (
	"fmt"
)

const top = 100000000

type gauss struct {
	r int
	i int
}

// https://projecteuler.net/problem=153
func main() {
	ans := 0
	z := gauss{1, 0}
	for realPart := 1; z.squareNorm() <= top; realPart++ {
		z.r = realPart
		for imPart := 0; z.squareNorm() <= top; imPart++ {
			z.i = imPart
			if greatestCommonDivisor(int64(z.i), int64(z.r)) != 1 {
				continue
			}
			delta := line(&z)
			ans += delta
		}
		//fmt.Printf("(%d) total:%d\t\t %d+%di\n", len(cheat), ans, z.r, z.i)
		z.i = 0
	}

	fmt.Println("153/ Investigating Gaussian Integers")
	fmt.Println(ans)
}

func (z *gauss) squareNorm() int {
	return z.r*z.r + z.i*z.i
}

var cheat = make(map[int]int)

func line(z *gauss) (total int) {
	mag := z.squareNorm()

	if ans, ok := cheat[top/mag]; ok {
		total = ans
	} else {
		for a := 1; mag*a <= top; a++ {
			for b := a; b*a*mag <= top; b++ {
				if b == a {
					total += a
				} else {
					total += (a + b)
				}
			}
		}
		cheat[top/mag] = total
	}

	total *= z.r

	if z.i > 0 {
		total *= 2
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
