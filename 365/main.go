package main

import (
	"fmt"
	"math"
	"math/big"
)

const (
	top    = 5000
	bottom = 1000
)

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

var (
	NCONST = IntExp(10, 18)
	KCONST = IntExp(10, 9)
)

// https://projecteuler.net/problem=365
func main() {
	PrimeCache(top)

	startP := PrimePi(bottom) + 1
	stopP := PrimePi(top)

	fmt.Println(Prime(startP), Prime(stopP))

	lookup := make([][2]int64, 0)
	for pindex := startP; pindex <= stopP; pindex++ {
		p := Prime(pindex)
		val := ChooseModP(NCONST, KCONST, p)

		lookup = append(lookup, [2]int64{p, val})
	}

	size := len(lookup)
	fmt.Println("built table with", size, "entries")

	total := big.NewInt(0)

	for i := 0; i < size; i++ {
		a1, p1 := lookup[i][1], lookup[i][0]

		for j := i + 1; j < size; j++ {
			a2, p2 := lookup[j][1], lookup[j][0]

			for k := j + 1; k < size; k++ {
				a3, p3 := lookup[k][1], lookup[k][0]

				a := []int64{a1, a2, a3}
				p := []int64{p1, p2, p3}

				crt := BigChineseRemainder(a, p)

				total.Add(total, crt)
			}
		}
		fmt.Println(1+i, "of", size, "(", total, ")")
	}

	fmt.Println("365/")
	fmt.Println(total)
}

func PrimePi(n int64) int64 {
	if n < 2 {
		return 0
	}

	if n == 2 {
		return 1
	}

	if n < primeTableLength && primepilist[n] != 0 {
		return primepilist[n]
	}

	var answer int64
	if IsPrime(n) {
		answer = 1 + PrimePi(n-1)
		if answer < primeTableLength {
			primeTable[answer] = n
		}
	} else {
		answer = PrimePi(n - 1)

	}

	if n < primeTableLength {
		primepilist[n] = answer
	}
	return answer
}

func PrimeCache(n int64) {
	composite := make([]bool, n)
	composite[0], composite[1] = true, true

	p := int64(2)

	for p < n {
		marker := p + p
		for ; marker < n; marker += p {
			composite[marker] = true
		}
		p++
		for ; p < n && composite[p]; p++ {
		}
	}

	m := 1
	for i := int64(0); i < n; i++ {
		if !composite[i] {
			primeTable[m] = i

			if int(i) < len(primepilist) {
				primepilist[i] = int64(m)
			}
			m++
		}
	}
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

func ChooseModP(n, k, p int64) int64 {
	ans := int64(1)

	//The case where Lucas' Theorem applies
	if n > p || k > p {
		for n > 0 || k > 0 {
			thisN, thisK := n%p, k%p
			n /= p
			k /= p
			if thisN < thisK {
				return 0
			}

			ans *= ChooseModP(thisN, thisK, p)
			ans %= p
		}
		return ans
	}

	//The case where Lucas' Theorem does not apply
	for i := int64(1); i <= k; i++ {
		ans *= n + 1 - i
		ans %= p
		ans *= InverseMod(i, p)
		ans %= p
	}
	return ans
}

func BigChineseRemainder(a, n []int64) *big.Int {
	N := big.NewInt(1)
	temp := big.NewInt(0)
	temp2 := big.NewInt(0)
	summand := big.NewInt(0)

	for _, en := range n {
		N.Mul(N, temp.SetInt64(en))
	}

	ans := big.NewInt(0)
	for i := range a {
		summand.SetInt64(a[i])
		summand.Mul(summand, N)
		summand.Div(summand, temp.SetInt64(n[i]))

		temp.SetInt64(n[i])
		summand.Mul(summand, BigInverseMod(temp2.Div(N, temp), temp))

		ans.Add(ans, summand)
	}

	ans.Mod(ans, N)

	return ans
}

func InverseMod(x, n int64) int64 {
	ans, _ := ExtendedEuclidean(x, n)

	ans %= n
	for ans < 0 {
		ans += n
	}

	return ans
}

func BigInverseMod(x, n *big.Int) *big.Int {
	ans, _ := BigExtendedEuclidean(x, n)
	return ans
}

func ExtendedEuclidean(a, b int64) (x, y int64) {
	x, lastx := int64(0), int64(1)
	y, lasty := int64(1), int64(0)
	for b != int64(0) {
		quotient := a / b
		a, b = b, a%b
		x, lastx = lastx-quotient*x, x
		y, lasty = lasty-quotient*y, y
	}
	return lastx, lasty
}

func BigExtendedEuclidean(a, b *big.Int) (x, y *big.Int) {
	x, lastx := big.NewInt(0), big.NewInt(1)
	y, lasty := big.NewInt(1), big.NewInt(0)

	zero := big.NewInt(0)
	quotient := big.NewInt(0)
	temp, temp2 := big.NewInt(0), big.NewInt(0)

	for b.Cmp(zero) != 0 {
		quotient.Div(a, b)

		temp.Set(b)
		b.Mod(a, b)
		a.Set(temp)

		temp.Set(x)
		x.Sub(lastx, temp2.Mul(quotient, x))
		lastx.Set(temp)

		temp.Set(y)
		y.Sub(lasty, temp2.Mul(quotient, y))
		lasty.Set(temp)

	}
	return lastx, lasty
}
