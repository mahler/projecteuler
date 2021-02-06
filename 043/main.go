package main

import (
	"fmt"
	"strconv"
)

// The number, 1406357289, is a 0 to 9 pandigital number because it is made up of
// each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string
// divisibility property.
//
// Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:
//
// d2d3d4=406 is divisible by 2
// d3d4d5=063 is divisible by 3
// d4d5d6=635 is divisible by 5
// d5d6d7=357 is divisible by 7
// d6d7d8=572 is divisible by 11
// d7d8d9=728 is divisible by 13
// d8d9d10=289 is divisible by 17
//
// Find the sum of all 0 to 9 pandigital numbers with this property.

var divisors []int
var tmp int
var tmp64 int64

func main() {
	var result int64
	divisors = []int{2, 3, 5, 7, 11, 13, 17}
	result = allPerms("", "0123456789")
	fmt.Println("043/ Find the sum of all 0 to 9 pandigital numbers with this property.")
	fmt.Println(result)
}

func allPerms(pre, s string) int64 {
	if len(s) == 0 {
		return checkTheProperty(pre)
	} else {
		var returnResult int64
		for i := 0; i < len(s); i++ {
			returnResult += allPerms(pre+s[i:i+1], s[0:i]+s[i+1:len(s)])
		}
		return returnResult
	}
}

func checkTheProperty(s string) int64 {
	if s[0] == '0' {
		return 0
	} else {
		for i := 1; i < 8; i++ {
			if tmp, _ = strconv.Atoi(s[i : i+3]); tmp%divisors[i-1] != 0 {
				return 0
			}
		}
	}
	tmp, _ = strconv.Atoi(s)
	return int64(tmp)
}
