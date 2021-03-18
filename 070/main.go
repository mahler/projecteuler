package main

import (
	"fmt"
	"math"
	"strconv"
)

var totientTable [totientTableLength]int64

const totientTableLength = 100000
const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

// Euler's Totient function, φ(n) [sometimes called the phi function], is used to determine the number
// of positive numbers less than or equal to n which are relatively prime to n. For example, as 1, 2, 4,
// 5, 7, and 8, are all less than nine and relatively prime to nine, φ(9)=6.
//
// The number 1 is considered to be relatively prime to every positive number, so φ(1)=1.
//
// Interestingly, φ(87109)=79180, and it can be seen that 87109 is a permutation of 79180.
//
// Find the value of n, 1 < n < 107, for which φ(n) is a permutation of n and the ratio n/φ(n) produces a minimum.
func main() {
	record := 100.
	var result int64
	for i := int64(0); i < 10000000; i++ {
		if ArePermutations(i, Totient(i)) {
			ratio := float64(i) / float64(Totient(i))
			if ratio < record {
				record = ratio
				// fmt.Println(i, record)
				result = i
			}
		}

	}

	fmt.Println("070/ Find the value of n, 1 < n < 107, for which φ(n) is a permutation of n and the ratio n/φ(n) produces a minimum.")
	fmt.Println(result)

}

func Totient(n int64) int64 {
	if n < 2 {
		return 0
	}
	if n < totientTableLength && totientTable[n] != 0 {
		return totientTable[n]
	}

	factors := Factor(n)

	if factors[0] == factors[len(factors)-1] {
		answer := IntExp(factors[0], int64(len(factors))) - IntExp(factors[0], int64(len(factors)-1))
		if n < totientTableLength {
			totientTable[n] = answer
		}
		return answer
	}

	for i := 0; i < len(factors); i++ {
		if factors[i] != factors[0] {
			split := IntExp(factors[0], int64(i))
			answer := Totient(split) * Totient(n/split)
			if n < totientTableLength {
				totientTable[n] = answer
			}
			return answer
		}
	}

	//bad things have happenned if we're here
	return 0

}

func ArePermutations(a int64, b int64) bool {
	A := strconv.FormatInt(a, 10)
	B := strconv.FormatInt(b, 10)

	length := len(A)
	list1 := make([]byte, length)
	list2 := make([]byte, length)

	if len(A) != len(B) {
		return false
	}

	for i := 0; i < length; i++ {
		list1[i] = A[i]
		list2[i] = B[i]
	}

	for i := 0; i < length; i++ {
		flag := false

		for j := 0; j < length; j++ {
			if flag == false && list1[i] == list2[j] {
				list2[j] = 0
				flag = true
			}

		}
		if flag == false {
			return false
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
