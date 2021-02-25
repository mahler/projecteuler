package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	tablesize = 10000000
	tailength = 1000000000
	zero      = .00001
)

var (
	fibTails [tablesize]int
	fibHeads [tablesize]science
)

type science struct {
	value  float64
	offset int
}

// https://projecteuler.net/problem=104
func main() {
	answer := 0
	for i := 10; !isPandigital(int64(fibHead(i).value)) || !isPandigital(int64(fibTail(i))); i++ {
		answer = i + 1
	}

	fmt.Println("104/ k value for result is:")
	fmt.Println(answer)
}

func fibTail(n int) int {

	if n <= 2 {
		return 1
	}

	if n < tablesize && fibTails[n] != 0 {
		return fibTails[n]
	}

	answer := (fibTail(n-1) + fibTail(n-2)) % tailength
	if n < tablesize {
		fibTails[n] = answer
	}
	return answer

}

func fibHead(n int) science {

	if n <= 2 {
		return science{1, 0}
	}

	if n < tablesize && fibHeads[n].value > zero {
		return fibHeads[n]
	}

	part1 := fibHead(n - 1)
	part2 := fibHead(n - 2)

	var answer science

	if part1.offset == part2.offset {
		answer.value = part1.value + part2.value
		answer.offset = part1.offset
	}

	if part1.offset > part2.offset {
		answer.value = part1.value + (part2.value / 10)
		answer.offset = part1.offset
	}

	if answer.value > tailength {
		answer.value = answer.value / 10
		answer.offset++
	}

	if n < tablesize {
		fibHeads[n] = answer
	}
	return answer

}

func isPandigital(n int64) bool {

	height := 1 + int64(math.Log10(float64(n)))

	output := int64(0)

	for i := int64(1); i < height+1; i++ {
		current := int64(1)
		for j := int64(1); j < i; j++ {
			current *= 10
		}
		output += (current * (i))
	}

	return output == SortInt(n)
}

func SortInt(input int64) int64 {

	swapped, _ := strconv.ParseInt(Sort(strconv.FormatInt(input, 10)), 10, 64)

	return swapped

}

func IsPandigital(n int64) bool {

	height := 1 + int64(math.Log10(float64(n)))

	output := int64(0)

	for i := int64(1); i < height+1; i++ {
		current := int64(1)
		for j := int64(1); j < i; j++ {
			current *= 10
		}
		output += (current * (i))
	}

	return output == SortInt(n)
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
