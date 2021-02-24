package main

import (
	"fmt"
)

const size = 7
const start = 19
const maxgap = 11

func main() {

	best := 99999
	var result []int

	for i := 0; i < int(IntExp(maxgap, size)); i++ {
		try := test(i)
		if total(try) < best && check2(try) && check1(try) {
			best = total(try)
			result = test(i)
		}
	}
	fmt.Println("103/ Given that A is an optimum special sum set for n = 7, find its set string")
	for _, r := range result {
		fmt.Print(r)
	}
	fmt.Println()
}

func binary(bin, index, siz int) []int {
	if siz == 0 {
		return []int{}
	}

	return append(binary(bin, index/bin, siz-1), index%bin+1)
}

func test(i int) []int {
	ret := binary(maxgap, i, size)
	ret[0] += start
	for i := 1; i < size; i++ {
		ret[i] += ret[i-1]
	}
	return ret
}

func check1(in []int) bool {
	set := make(map[int]bool)
	for i := 1; i < int(Exp2(size)); i++ {
		memb := binary(2, i, size)
		sum := 0
		for j := 0; j < size; j++ {
			if memb[j] == 1 {
				sum += in[j]
			}
		}
		if set[sum] {
			return false
		}
		set[sum] = true
	}

	return true
}

func check2(in []int) bool {
	sum1, sum2 := in[0], 0
	for i := 0; i < size/2; i++ {
		sum1 += in[i+1]
		sum2 += in[size-i-1]
		if sum1 < sum2 {
			return false
		}
	}

	return true
}

func total(in []int) (sum int) {
	for _, x := range in {
		sum += x
	}
	return
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

func Exp2(n int) int64 {
	answer := int64(1)
	for i := 0; i < n; i++ {
		answer *= 2
	}
	return answer
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
