package main

import (
	"fmt"
	"math"
)

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64
var borMemo map[[2]int]uint64

func main() {
	borMemo = make(map[[2]int]uint64)

	sum := uint64(0)
	tip, top := 47, 43

	for i := 1; i <= tip; i++ {
		for j := 1; j <= top; j++ {
			sum += blend(i, j)
		}
	}

	fmt.Println("147/ How many different rectangles could be situated within 47x43 and smaller grids?")
	fmt.Println(sum)
}

//returns the number of subrectangles in a x b grid
//with the boring orientation
func boring(a, b int) uint64 {

	if ans, ok := borMemo[[2]int{a, b}]; ok {
		return ans
	}

	if a > 1 {
		//we decompose into configurations which avoid the top row
		//those which avoid the bottom row, and those of full height
		borMemo[[2]int{a, b}] = 2*boring(a-1, b) - boring(a-2, b) + boring(1, b)
		return boring(a, b)
	}

	if a == 1 && b > 1 {
		borMemo[[2]int{a, b}] = uint64(b) + boring(a, b-1)
		return boring(a, b)
	}

	if a <= 0 || b <= 0 {
		return 0
	}

	//a==b==1
	return 1

}

func sort(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a, b
}

func predict(a, b int) uint64 {
	b, a = sort(a, b)

	if a == 1 {
		return uint64(b - 1)
	}

	offset := 0

	factors := Factors(int64(a))

	if factors[0][0] > 3 {
		offset = -1
	}

	answer := int64(offset - (a / 2) + (a*a)/6 - (2*a*a*a*a)/3 - (a*b)/3 + (4*a*a*a*b)/3)

	if answer < 0 {
		panic("negative answer is bad")
	}

	return uint64(answer)
}

func blend(a, b int) uint64 {

	return boring(a, b) + predict(a, b)
}

func Factor(n int64) []int64 {
	var answer = []int64{}

	current := n

	i := int64(1)
	for !IsPrime(current) {
		if current%Prime(i) == 0 {
			answer = append(answer, Prime(i))
			current = current / Prime(i)
			i = 0
		}
		i++
	}

	answer = append(answer, current)

	return answer
}

func Factors(n int64) [][2]int64 {
	factorList := Factor(n)
	factors := [][2]int64{[2]int64{factorList[0], 1}}
	for i := 1; i < len(factorList); i++ {
		if factorList[i] == factors[len(factors)-1][0] {
			factors[len(factors)-1][1]++
		} else {
			factors = append(factors, [2]int64{factorList[i], 1})
		}
	}
	return factors
}

func IsPrime(n int64) bool {

	if n == 1 {
		return false
	}

	end := int64(math.Sqrt(float64(n)))

	//If we start computing beyond the table, this is stupid
	for i := int64(1); Prime(i) <= end && i < primeTableLength; i++ {
		if n%Prime(i) == 0 {
			return false
		}
	}

	//If we need to pass the end of the prime table brute force
	if end > lastPrime {
		for i := int64(lastPrime); i <= end; i++ {
			if n%i == 0 {
				return false
			}
		}

	}

	return true
}

func Prime(n int64) int64 {

	if n < 1 {
		return 0
	}

	primeTable[1] = 2
	primeTable[2] = 3

	if n < primeTableLength && primeTable[n] != 0 {
		return primeTable[n]
	}

	i := Prime(n-1) + 1

	for !IsPrime(i) {
		i++
	}

	if i < primeTableLength {
		primepilist[i] = n
	}

	if n < primeTableLength {
		primeTable[n] = i
	}
	return i
}
