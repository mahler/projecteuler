package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// A common security method used for online banking is to ask the user for three random
// characters from a passcode. For example, if the passcode was 531278, they may ask
// for the 2nd, 3rd, and 5th characters; the expected reply would be: 317.
//
// The text file, keylog.txt, contains fifty successful login attempts.
//
// Given that the three characters are always asked for in order, analyse the file so as to determine the shortest possible secret passcode of unknown length.

func main() {
	var passcode string
	data, err := os.ReadFile("p079_keylog.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	fileRows := strings.Split(string(data), "\n")

	var code [][]rune

	for _, fileRow := range fileRows {
		code = append(code, []rune(fileRow))
	}

	var link [10][10]bool

	for i := 0; i < 50; i++ {
		link[code[i][0]-'0'][code[i][1]-'0'] = true
		link[code[i][1]-'0'][code[i][2]-'0'] = true
	}

	var table [10]bool
	var inDeg, outDeg [10]int

	for i := 0; i < 10; i++ {
		table[i] = false
		inDeg[i] = 0
		outDeg[i] = 0
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if link[i][j] {
				table[i] = true
				table[j] = true
				outDeg[i]++
				inDeg[j]++
			}
		}
	}

	left := 0
	for _, x := range table {
		if x {
			left++
		}
	}

	for left > 0 {
		curr := 0
		for curr < 10 {
			if table[curr] && inDeg[curr] == 0 {
				passcode += fmt.Sprintf("%v", curr)
				table[curr] = false
				left--

				for i := 0; i < 10; i++ {
					if link[i][curr] {
						outDeg[i]--
						inDeg[curr]--
					}
					if link[curr][i] {
						outDeg[curr]--
						inDeg[i]--
					}
				}

				break
			} else {
				curr++
			}
		}
	}

	fmt.Println("079/ the shortest possible secret passcode of unknown length?")
	fmt.Println(passcode)
}

func IsValid(x, y []rune) bool {
	indexX, indexY := 0, 0

	for indexX < len(x) && indexY < len(y) {
		if x[indexX] == y[indexY] {
			indexY++
		}
		indexX++
	}

	return indexY == len(y)
}
