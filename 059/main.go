package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var ciphertext []int

func main() {
	ciphertext = make([]int, 0)
	fileContent, err := os.ReadFile("p059_cipher.txt")
	if err != nil {
		log.Fatal("File reading error", err)
		return
	}

	in := string(fileContent)
	split := strings.Split(in, ",")

	for _, number := range split {
		integer, _ := strconv.Atoi(number)
		ciphertext = append(ciphertext, integer)
	}

	max := 0
	key := make([]int, 3)

	for i := 97; i <= 122; i++ {
		for j := 97; j <= 122; j++ {

			for k := 97; k <= 122; k++ {

				try := []int{i, j, k}

				cleartext := decode(try)

				score := 0

				for _, letter := range cleartext {

					if letter == 32 {
						score++
					}
				}

				if score > max {
					max = score
					key = []int{i, j, k}
				}
			}
		}
	}

	fmt.Print("Key: ", key, " = ")
	for _, letter := range key {
		fmt.Print(string(letter))
	}
	fmt.Print("\n")

	cleartext := decode(key)

	var sum int
	for _, letter := range cleartext {
		fmt.Print(string(letter))
		sum += letter
	}

	fmt.Printf("\n%d\n", sum)
}

func decode(password []int) (cleartext []int) {
	cleartext = make([]int, len(ciphertext))
	for i, letter := range ciphertext {
		cleartext[i] = letter ^ password[i%len(password)]
	}
	return
}
