package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
)

// Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.
//
// [puzzle.txt]

func main() {
	data, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileContent := strings.Split(string(data), "\n")

	bigSum := new(big.Int)
	bigSum.SetString("0", 10)

	for _, value := range fileContent {
		bigNum, _ := new(big.Int).SetString(value, 10)
		bigSum.Add(bigSum, bigNum)
	}

	fmt.Println("Full sum:")
	fmt.Println(bigSum)
	fmt.Println()
	strBigSum := fmt.Sprintf("%s", bigSum)
	fmt.Println("First 10 digits:")
	fmt.Println(strBigSum[0:10])
}
