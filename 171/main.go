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

// For a positive integer n, let f(n) be the sum of the squares of the digits (in base 10) of n, e.g.
//
// f(3) = 3^2 = 9,
// f(25) = 2^2 + 5^2 = 4 + 25 = 29,
// f(442) = 42 + 42 + 22 = 16 + 16 + 4 = 36
//
// Find the last nine digits of the sum of all n, 0 < n < 10^20, such that f(n) is a perfect square.
func main() {

	hittable := make(map[int]bool)
	ans := int64(0)

	for i := 1; i*i <= top*9*9; i++ {
		hittable[i*i] = true
	}

	for i := range hittable {
		ways, _ := enumerate(9, top, i)
		for i := 0; i < len(ways); i++ {
			ans += process(ways[i])
			ans %= mod
		}
	}

	fmt.Println("171/ Find the last nine digits of the sum of all n, 0 < n < 10^20, such that f(n) is a perfect square")
	fmt.Println(ans)
}

func enumerate(max, lth, target int) (ans [][]int, ok bool) {
	if target < 0 {
		return nil, false
	}

	if max*max*lth < target {
		return nil, false
	}

	if lth == 1 {
		if IsSquare(int64(target)) {
			sqrt := int(math.Sqrt(float64(target)))
			return [][]int{{sqrt}}, true
		} else {
			return nil, false
		}
	}

	ans = make([][]int, 0)

	for i := 0; i <= max && i*i <= target; i++ {
		tail, works := enumerate(i, lth-1, target-(i*i))
		if works {
			ok = true
			ans = append(ans, paint(tail, i)...)
		}
	}

	return
}

func paint(topaint [][]int, color int) [][]int {
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

func process(way []int) int64 {
	comp := compress(way)
	size := int64(len(way))
	choices := int64(1)

	for i := 0; i < len(comp); i++ {
		choices *= Choose(size, int64(comp[i][1]))
		size -= int64(comp[i][1])
	}

	delta := int64(0)
	for i := 0; i < len(comp); i++ {
		delta += int64(comp[i][1]) * int64(comp[i][0])
		delta %= mod
	}

	delta *= choices
	delta /= int64(len(way))
	delta %= mod

	delta = (delta * ones)
	delta %= mod

	return delta
}

func IsSquare(x int64) bool {
	return int64(math.Sqrt(float64(x)))*int64(math.Sqrt(float64(x))) == x
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
