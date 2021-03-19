package main

import (
	"fmt"
	"strconv"
)

var memo map[string]int

// A particular school offers cash rewards to children with good attendance and punctuality. If they are absent for three consecutive days or late on more than one occasion then they forfeit their prize.
//
// During an n-day period a trinary string is formed for each child consisting of L's (late), O's (on time), and A's (absent).
//
// Although there are eighty-one trinary strings for a 4-day period that can be formed, exactly forty-three strings would lead to a prize:
//
// OOOO OOOA OOOL OOAO OOAA OOAL OOLO OOLA OAOO OAOA
// OAOL OAAO OAAL OALO OALA OLOO OLOA OLAO OLAA AOOO
// AOOA AOOL AOAO AOAA AOAL AOLO AOLA AAOO AAOA AAOL
// AALO AALA ALOO ALOA ALAO ALAA LOOO LOOA LOAO LOAA
// LAOO LAOA LAAO
//
// How many "prize" strings exist over a 30-day period?
func main() {
	memo = make(map[string]int)
	total := 30
	fmt.Println("191/ How many prize strings exist over a 30-day period?")
	fmt.Println(numb(total, "OO", false) + numb(total, "OA", false) + numb(total, "OL", false) + numb(total, "AO", false) + numb(total, "AA", false) + numb(total, "AL", false) + numb(total, "LO", false) + numb(total, "LA", false))
}

func numb(length int, leading2 string, hasbeenlate bool) int {

	if length == 2 {
		if leading2 == "AA" {
			return 1
		}
		if leading2 == "AL" {
			if hasbeenlate {
				return 0
			}
			return 1
		}
		if leading2 == "AO" {
			return 1
		}
		if leading2 == "LA" {
			if hasbeenlate {
				return 0
			}
			return 1
		}
		if leading2 == "LL" {
			return 0
		}
		if leading2 == "LO" {
			if hasbeenlate {
				return 0
			}
			return 1
		}
		if leading2 == "OA" {
			return 1

		}
		if leading2 == "OL" {
			if hasbeenlate {
				return 0
			}
			return 1
		}
		if leading2 == "OO" {
			return 1
		}
	}

	if answer, ok := memo[keygen(length, leading2, hasbeenlate)]; ok {
		return answer
	}

	a := leading2[0:1]
	b := leading2[1:2]

	if (a == "L" || b == "L") && hasbeenlate {
		return 0
	}

	answer := 0

	nextlate := hasbeenlate
	if a == "L" {
		nextlate = true
	}

	if a == b && a == "A" {

	} else {
		answer += numb(length-1, b+"A", nextlate)
	}

	answer += numb(length-1, b+"L", nextlate)
	answer += numb(length-1, b+"O", nextlate)

	memo[keygen(length, leading2, hasbeenlate)] = answer

	return answer

}

func keygen(n int, s string, b bool) string {
	answer := ""
	answer += strconv.Itoa(n)
	answer += s
	if b {
		answer += "T"
	} else {
		answer += "F"
	}
	return answer
}
