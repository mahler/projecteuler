package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
//
// What is the sum of the digits of the number 2^1000
func main() {
	var i, e = big.NewInt(2), big.NewInt(1000)
	i.Exp(i, e, nil)
	strRes := fmt.Sprint(i)

	var sum int
	for x := 0; x < len(strRes); x++ {
		i, _ := strconv.Atoi(strRes[x : x+1])
		sum += i
	}

	fmt.Println("016: Sum of digits:", sum)
}
