package main

import (
	"fmt"
	"math"
	"strconv"
)

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

const D = 10

// See https://projecteuler.net/problem=111
func main() {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	total := int64(0)

	for _, char := range digits {

		N := 0
		S := int64(0)
		m := 0
		for m = 0; N == 0; m++ {

			N, S = test(char, D, m)

		}

		total += S

	}

	fmt.Println("111/ Find the sum of all S(10, d)")
	fmt.Println(total)
}

func test(repeat string, n, d int) (N int, S int64) {

	for i := 0; int64(i) < Choose(int64(n), int64(d)); i++ {
		indices := SplitSeq(d, i)

		for j := 0; int64(j) < IntExp(10, int64(d)); j++ {
			insertstring := strconv.Itoa(j)
			for len(insertstring) < d {
				insertstring = "0" + insertstring
			}

			merged := ""
			current := 0
			for index := 0; index < n; index++ {
				if current < d && index == indices[d-current-1] {
					merged += insertstring[current : current+1]
					current++
				} else {
					merged += repeat
				}
			}

			mergedint, _ := strconv.ParseInt(merged, 10, 64)

			//exclude leading zeroes
			if mergedint > IntExp(10, int64(n)-1) {

				if isPrime(mergedint) {
					//fmt.Println(mergedint)
					N++
					S += mergedint
				}
			}
		}
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
