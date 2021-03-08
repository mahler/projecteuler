package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	total := 0
	i := 1
	for 100*total < 99*i {
		i++
		if isBouncy(i) {
			total++
		}
	}

	fmt.Println("Ratio:", total, "/", i, " = 99%")
	fmt.Println()
	fmt.Println("112/ Find the least number for which the proportion of bouncy numbers is exactly 99%")
	fmt.Println(i)
}

func isBouncy(n int) bool {
	x := int(SortInt(int64(n)))
	if x == n {
		return false
	}

	y := int(IntReverse(int64(x)))
	if y == n {
		return false
	}

	return true
}

func SortInt(input int64) int64 {
	swapped, _ := strconv.ParseInt(Sort(strconv.FormatInt(input, 10)), 10, 64)
	return swapped
}

func IntReverse(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	s = StringReverse(s)
	x, _ := strconv.ParseInt(s, 10, 64)
	return x
}

func Sort(word string) string {
	wordtable := strings.Split(word, "")
	for j := 0; j < len(word); j++ {

		for i := 0; i < len(word)-1; i++ {
			if wordtable[i] < wordtable[i+1] {
				temp := wordtable[i]
				wordtable[i] = wordtable[i+1]
				wordtable[i+1] = temp
			}
		}
	}
	return strings.Join(wordtable, "")
}

func StringReverse(a string) string {
	b := ""
	for i := len(a) - 1; i >= 0; i-- {
		b += string(a[i])
	}
	return b
}
