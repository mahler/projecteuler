package main

import (
	"fmt"
	"strconv"
	"strings"
)

// The cube, 41063625 (345^3), can be permuted to produce two other cubes: 56623104 (384^3) and 66430125 (4053). In fact, 41063625 is the smallest cube which has exactly three permutations of its digits which are also cube.
//
//Find the smallest cube for which exactly five permutations of its digits are cube.
func main() {
	var cubes = map[int64]int{}

	var x, i int64

	x, i = 1, 1
	for ; cubes[x] < 5; i++ {
		x = SortInt(i * i * i)
		cubes[x]++
	}

	x, i = 1, 1
	for ; cubes[x] != 5; i++ {
		x = SortInt(i * i * i)
	}

	i = i - 1
	fmt.Println("062/ Find the smallest cube for which exactly five permutations of its digits are cube.")
	fmt.Println(i * i * i)
}

func SortInt(input int64) int64 {
	swapped, _ := strconv.ParseInt(Sort(strconv.FormatInt(input, 10)), 10, 64)
	return swapped
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
