package main

import (
	"fmt"
	"math"
)

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629
const top = 18

// How many 18-digit numbers n (without leading zeros) are there such that no digit occurs more than three times in n?
func main() {
	total := int64(0)
	vectors := ways(10, top, 3)
	for _, vector := range vectors {
		total += distribute(vector)
	}

	fmt.Println("172/ How many 18-digit numbers n (without leading zeros) are there such that no digit occurs more than three times in n?")
	fmt.Println(total)
}

//the ways in which we can add up to goal, with descending sequences of n
//integers, whose values are at most, most
func ways(n, goal, most int) [][]int {
	if goal < 0 || most*n < goal {
		return nil
	}

	if n == 1 {
		return [][]int{{goal}}
	}

	accumulate := make([][]int, 0)

	for i := 0; i <= most; i++ {
		if below := ways(n-1, goal-i, i); below != nil {
			accumulate = append(accumulate, paint(i, below)...)
		}
	}

	return accumulate
}

func paint(color int, topaint [][]int) [][]int {
	ans := make([][]int, 0)
	for _, strip := range topaint {
		ans = append(ans, append(strip, color))
	}
	return ans
}

//returns compressed representation of slice
func compress(input []int) [][2]int {
	ans := [][2]int{{input[0], 1}}
	pointer := 0
	for i := 1; i < len(input); i++ {
		if ans[pointer][0] == input[i] {
			ans[pointer][1]++
		} else {
			ans = append(ans, [2]int{input[i], 1})
			pointer++
		}
	}
	return ans
}

func distribute(state []int) (ans int64) {
	types := len(compress(state))

	for zeroPos := 0; zeroPos < types; zeroPos++ {
		thisAns := int64(1)
		openPlaces := top
		flat := compress(state)

		thisAns *= Choose(int64(openPlaces-1), int64(flat[zeroPos][0]))
		openPlaces -= flat[zeroPos][0]
		flat[zeroPos][1]--

		openDigits := 9

		for iPos := 0; iPos < types; iPos++ {

			thisAns *= Choose(int64(openDigits), int64(flat[iPos][1]))
			openDigits -= flat[iPos][1]

			for i := 0; i < flat[iPos][1]; i++ {
				thisAns *= Choose(int64(openPlaces), int64(flat[iPos][0]))
				openPlaces -= flat[iPos][0]
			}
		}
		ans += thisAns
	}
	return
}

func Choose(N, K int64) int64 {
	factors := make(map[int64]int64)

	if K == 0 || N == K || N <= 1 {
		return 1
	}

	if N < K {
		return 0
	}

	for n := N; n > N-K; n-- {
		nfactors := findFactors(n)
		for i := 0; i < len(nfactors); i++ {
			factors[nfactors[i][0]] += nfactors[i][1]
		}
	}

	for k := K; k >= 2; k-- {

		kfactors := findFactors(k)
		for i := 0; i < len(kfactors); i++ {
			factors[kfactors[i][0]] -= kfactors[i][1]
		}

	}

	answer := int64(1)

	for prime, multiplicity := range factors {

		for i := int64(0); i < multiplicity; i++ {
			answer *= prime
		}

	}
	return answer
}

func SplitSeq(K, N int) (a []int) {

	indices := make([]int, 0)

	for k := K; k >= 1; k-- {

		n := k - 1

		if Choose(int64(n), int64(k)) <= int64(N) {
			for ; Choose(int64(n), int64(k)) <= int64(N); n++ {

			}
			n--
		}

		indices = append(indices, n)

		N = N - int(Choose(int64(n), int64(k)))
	}

	return indices
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

func isPrime(n int64) bool {

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

	for !isPrime(i) {
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

func findFactors(n int64) [][2]int64 {
	factorList := factor(n)
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

func factor(n int64) []int64 {
	var answer = []int64{}

	current := n

	i := int64(1)
	for !isPrime(current) {
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
