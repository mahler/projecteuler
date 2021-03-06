package main

import (
	"fmt"
	"os"
)

// For a number written in Roman numerals to be considered valid there are basic rules which must be followed. Even though the rules allow
// some numbers to be expressed in more than one way there is always a "best" way of writing a particular number.
//
// For example, it would appear that there are at least six ways of writing the number sixteen:
//
// IIIIIIIIIIIIIIII
// VIIIIIIIIIII
// VVIIIIII
// XIIIIII
// VVVI
// XVI
//
// However, according to the rules only XIIIIII and XVI are valid, and the last example is considered to be the most efficient, as it uses the least number of numerals.
//
// The 11K text file, roman.txt (right click and 'Save Link/Target As...'), contains one thousand numbers written in valid, but not
// necessarily minimal, Roman numerals; see About... Roman Numerals for the definitive rules for this problem.
//
// Find the number of characters saved by writing each of these in their minimal form.
//
// Note: You can assume that all the Roman numerals in the file contain no more than four consecutive identical units.
func main() {

	data := fileImport("p089_roman.txt")

	blerg := 0
	for _, numeral := range data {

		blerg += len(numeral) - best(evalNumeral(numeral))
	}

	fmt.Println("089/ Find the number of characters saved by writing each of these in their minimal form")
	fmt.Println(blerg)

}

func evalNumeral(numeral string) int {
	if len(numeral) == 0 {
		return 0
	}
	if len(numeral) == 1 {
		if numeral == "I" {
			return 1
		}
		if numeral == "V" {
			return 5
		}
		if numeral == "X" {
			return 10
		}
		if numeral == "L" {
			return 50
		}
		if numeral == "C" {
			return 100
		}
		if numeral == "D" {
			return 500
		}
		if numeral == "M" {
			return 1000
		}
	}

	a := numeral[len(numeral)-2 : len(numeral)-1]
	b := numeral[len(numeral)-1:]
	rest := numeral[:len(numeral)-2]

	A := evalNumeral(a)
	B := evalNumeral(b)

	if A < B {
		return evalNumeral(rest) + B - A
	}

	return B + evalNumeral(rest+a)
}

func best(n int) int {

	char1 := []int{1, 5, 10, 50, 100, 500, 1000}
	char2 := []int{4, 9, 40, 90, 400, 900}

	if n == 0 {
		return 0
	}

	a, b := 1, 0

	for i := 0; i < len(char1); i++ {
		if char1[i] <= n {
			a = char1[i]
		}
	}
	for i := 0; i < len(char2); i++ {
		if char2[i] <= n {
			b = char2[i]
		}
	}

	if b != 0 && b > a {
		return 2 + best(n-b)
	}

	return 1 + best(n-a)
}

func fileImport(filename string) []string {
	b, _ := os.ReadFile(filename)

	var output []string
	var currentline string

	for _, char := range b {
		if char == 10 {
			output = append(output, currentline)
			currentline = ""
		} else {
			currentline += string(char)
		}
	}

	if currentline != "" {
		output = append(output, currentline)
	}

	return output

}
