package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("p099_base_exp.txt")
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

	fmt.Println(ans, value)
}
