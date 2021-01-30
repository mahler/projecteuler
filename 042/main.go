package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten
// triangle numbers are:
//
// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
//
// By converting each letter in a word to a number corresponding to its alphabetical position
// and adding these values we form a word value. For example, the word value for SKY is
// 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.
//
// Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly
// two-thousand common English words, how many are triangle words?
var triangles []int

func main() {
	//read names
	fileContent, _ := ioutil.ReadFile("p042_words.txt")

	fileStr := string(fileContent)
	words := strings.Split(fileStr, ",")
	for i, word := range words {
		words[i] = word[1 : len(word)-1] // chop quotes from word
	}

	//create an array of triangle numbers
	triangles = make([]int, 20)
	for i := range triangles {
		triangles[i] = (i * (i + 1)) / 2
	}

	//iterate over names to check
	wordCount := 0
	for i := range words {
		if isTriangle(alphabeticValue(words[i])) {
			wordCount++
		}
	}
	fmt.Println("042/ how many are triangle words?")
	fmt.Println(wordCount)
}

func isTriangle(in int) bool {
	for i := range triangles {
		if triangles[i] == in {
			return true
		}
	}
	return false
}

func alphabeticValue(in string) int {
	var value int
	for i := range in {
		value += int(in[i]) - int('A') + 1
	}
	return value
}
