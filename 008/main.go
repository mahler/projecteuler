package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.
// See [puzzle.txt]
// Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?
//
// NOTE: int64 is needed for product calculations to avoid overrun of int32 range.

func main() {
	adjacent := 13

	data, err := ioutil.ReadFile("puzzle.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileContent := strings.Split(string(data), "\n")

	var numbers []int64

	for _, value := range fileContent {
		for col := 0; col < len(value); col++ {
			intVal, _ := strconv.Atoi(value[col : col+1])
			numbers = append(numbers, int64(intVal))
		}
	}

	// fmt.Println(len(numbers))
	var maxProduct int64
	var maxIndex int
	for test := 0; test < len(numbers)-adjacent; test++ {
		product := int64(1)
		for i := 0; i < adjacent; i++ {
			product *= numbers[test+i]
		}

		if product > maxProduct {
			maxProduct = product
			maxIndex = test
		}
	}

	for i := 0; i < adjacent; i++ {
		fmt.Print(numbers[maxIndex+i])
		if i != adjacent-1 {
			fmt.Print(" X ")
		}
	}
	fmt.Print(" = ")
	fmt.Println(maxProduct)
}
