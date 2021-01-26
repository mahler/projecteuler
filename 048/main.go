package main

import (
	"fmt"
	"math/big"
)

// The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.
//
// Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.

func main() {
	var result big.Int
	for i := 1; i <= 1000; i++ {
		tmp := result
		result.Add(&tmp, powerToSelf(i))
		//fmt.Println(result.String())
	}
	strResult := result.String()

	lenght := len(strResult)
	fmt.Println(strResult[lenght-10 : lenght])

}

func powerToSelf(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	var result big.Int
	result.Exp(big.NewInt(int64(n)), big.NewInt(int64(n)), nil)
	return &result
}
