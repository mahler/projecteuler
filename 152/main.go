package main

import (
	"fmt"
	"math"
)

const top int64 = 80

const primeTableLength = 50000000 //+1!!!!!!!!!!!
const lastPrime = 982451629

var primeTable [primeTableLength]int64
var primepilist [primeTableLength]int64

func main() {
	topPrime := PrimePi(top)
	primeBasis := make([]int64, topPrime+1)

	for i := int64(2); i <= top; i++ {
		repr := rep(i * i)

		for _, frac := range repr {
			den := frac[1]
			pType := Factor(den)[0]
			pDex := PrimePi(pType)

			if den > primeBasis[pDex] {
				primeBasis[pDex] = den
			}
		}

	}

	canonical := make([][]int64, top+1)

	for i := int64(2); i <= top; i++ {
		standard := make([]int64, topPrime+1)

		repr := rep(i * i)

		for _, frac := range repr {
			den := frac[1]
			pType := Factor(den)[0]
			pDex := PrimePi(pType)
			offset := primeBasis[pDex] / den
			standard[pDex] = offset * frac[0]
		}

		canonical[i] = standard

	}

	predicates := make([][]int64, 1)
	predicates[0] = make([]int64, top+1)

	place := topPrime

	for place = topPrime; place >= 1; place-- {
		crinkTab := make([][2]int64, 0)

		for i, rep := range canonical {
			if len(rep) >= int(place+1) && rep[place] != 0 {
				modularity := rep[place] % primeBasis[place]
				crinkTab = append(crinkTab, [2]int64{int64(i), modularity})
			}
		}

		newPreds := make([][]int64, 0)

		for _, pred := range predicates {
			target := int64(0)
			if place == 1 {
				target = primeBasis[1] / 2
			}
			newPreds = append(newPreds, compute(pred, crinkTab, primeBasis[place], target)...)
		}

		predicates = newPreds

	}

	fmt.Println("152/ How many ways are there to write the number 1/2 as a sum of inverse squares using distinct integers between 2 and 80 inclusive?")
	fmt.Println(len(predicates))
}

func rep(x int64) [][]int64 {
	facs := findFactors(x)
	d1 := IntExp(facs[0][0], facs[0][1])

	if len(facs) == 1 {
		return [][]int64{{1, d1}}
	}

	d2 := (x) / d1
	num1, num2 := ExtendedEuclidean(d2, d1)
	up := times(rep(d2), num2)
	return append(up, []int64{num1, d1})
}

//helper function for rep
func times(fracs [][]int64, mult int64) [][]int64 {
	for i, val := range fracs {
		fracs[i][0] = val[0] * mult
	}
	return fracs
}

func compute(given []int64, pins [][2]int64, modulus int64, target int64) [][]int64 {
	concern := make([][2]int64, 0)
	for _, pin := range pins {
		ind, modifier := pin[0], pin[1]

		if given[ind] == 0 {
			concern = append(concern, pin)
		}
		if given[ind] == 1 {
			target -= modifier
		}

	}

	lambdas := make([]int64, len(concern))
	for i := 0; i < len(concern); i++ {
		lambdas[i] = concern[i][1]
	}

	results := dumbNumerate(lambdas, target, modulus)

	adjusted := make([][]int64, len(results))
	for i := 0; i < len(adjusted); i++ {
		adjusted[i] = make([]int64, len(given))
		copy(adjusted[i], given)

		for j, tuple := range concern {
			if results[i][j] {
				adjusted[i][tuple[0]] = 1
			} else {
				adjusted[i][tuple[0]] = -1
			}
		}
	}

	return adjusted
}

func dumbNumerate(vals []int64, target, modulus int64) (ans [][]bool) {
	ans = make([][]bool, 0)

	for state := 0; state < 1<<uint(len(vals)); state++ {

		crunch := int64(0)

		carrot := recover(state, len(vals))

		for i, on := range carrot {

			if on {
				crunch += vals[i]
			}

		}

		if (modulus+(crunch%modulus))%modulus == (modulus+(target%modulus))%modulus {
			ans = append(ans, carrot)
		}
	}

	return
}

func recover(ind, lent int) []bool {
	ans := make([]bool, lent)
	for i := 0; i < lent; i++ {
		if ind%2 == 1 {
			ans[i] = true
		}
		ind /= 2

	}
	return ans
}

func compatible(a, b []int64) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == 0 || b[i] == 0 {
			continue
		}

		if a[i] != b[i] {
			return false
		}
	}
	return true
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

func findFactors(n int64) [][2]int64 {
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
