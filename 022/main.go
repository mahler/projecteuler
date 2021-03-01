package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Using p022_names.txt, a 46K text file containing over five-thousand first names,
// begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value
// by its alphabetical position in the list to obtain a name score.
//
// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th
// name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.
//
// What is the total of all the name scores in the file?

func main() {
	data, err := os.ReadFile("p022_names.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}
	fileContent := strings.Split(string(data), ",")
	// Sort list alphabetically
	sort.Strings(fileContent)

	var nameSum int
	for number, aName := range fileContent {
		aName := aName[1 : len(aName)-1]
		nameSum += nameValue(aName) * (number + 1)

		//	fmt.Println(number+1, aName, nameValue(aName), "- Score:", nameValue(aName)*(number+1))

	}
	fmt.Println("022/ What is the total of all the name scores in the file?")
	fmt.Println(nameSum)
}

func nameValue(s string) int {
	var sum int

	r := []rune(s)
	for _, runeVal := range r {
		sum += int(runeVal) - 64
	}
	return sum
}
