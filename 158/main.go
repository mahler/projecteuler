package main

import (
	"fmt"
	"math"
)

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

const (
	top              int = 20
	mod                  = 1000000000 //10^9
	ones                 = 111111111  // nine 1s
	primeTableLength     = 50000000   //+1!!!!!!!!!!!
	lastPrime            = 982451629
)

// Taking three different letters from the 26 letters of the alphabet, character strings of
// length three can be formed. Examples are 'abc', 'hat' and 'zyx'.
//
// When we study these three examples we see that for 'abc' two characters come lexicographically
// after its neighbour to the left. For 'hat' there is exactly one character that comes lexicographically
// after its neighbour to the left. For 'zyx' there are zero characters that come lexicographically after
// its neighbour to the left. In all there are 10400 strings of length 3 for which exactly one character
// comes lexicographically after its neighbour to the left.
//
// We now consider strings of n ≤ 26 different characters from the alphabet.
// For every n, p(n) is the number of strings of length n for which exactly one character comes
// lexicographically after its neighbour to the left.
//
// What is the maximum value of p(n)?
func main() {
	record := int64(-1)

	for i := int64(1); i < 26; i++ {
		if ways(i) > record {
			record = ways(i)
		} else {
			break
		}
	}

	fmt.Println("158/ What is the maximum value of p(n)?")
	fmt.Println(record)
}

func ways(k int64) int64 {
	return (Exp2(int(k)) - 1 - k) * Choose(26, k)
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

func Exp2(n int) int64 {
	answer := int64(1)
	for i := 0; i < n; i++ {
		answer *= 2
	}
	return answer
}
