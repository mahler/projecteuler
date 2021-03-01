package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Comparing two numbers written in index form like 211 and 37 is not difficult, as any calculator would confirm that 2^11 = 2048 < 3^7 = 2187.
//
// However, confirming that 632382^518061 > 519432^525806 would be much more difficult, as both numbers contain over three million digits.
//
// Using base_exp.txt, a 22K text file containing one thousand lines with a base/exponent pair on each line, determine which line number
// has the greatest numerical value.

func main() {
	data, err := os.ReadFile("p099_base_exp.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}
	fileContent := strings.Split(string(data), "\n")

	ans := 0
	value := float64(0)

	for i, val := range fileContent {
		s := strings.Split(val, ",")
		base, _ := strconv.Atoi(s[0])
		exp, _ := strconv.Atoi(s[1])
		nextValue := math.Log10(float64(base)) * float64(exp)
		if nextValue > value {
			ans = i + 1
			value = nextValue
		}
	}

	fmt.Println("099/ Determine which line number has the greatest numerical value")
	fmt.Println(ans, value)
}
