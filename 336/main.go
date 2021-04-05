package main

import (
	"fmt"
	"sort"
)

const top = 11
const which = 2011

type trains struct {
	contents [][]int
}

func (t trains) Len() int      { return len(t.contents) }
func (t trains) Swap(i, j int) { t.contents[i], t.contents[j] = t.contents[j], t.contents[i] }
func (t trains) Less(i, j int) bool {
	for a := 0; a < len(t.contents[i]); a++ {
		if t.contents[i][a] > t.contents[j][a] {
			return false
		} else if t.contents[i][a] < t.contents[j][a] {
			return true
		}
	}
	return true
}

// https://projecteuler.net/problem=336
func main() {
	begin := [][]int{{3, 0, 2, 1}, {3, 1, 0, 2}}
	var next [][]int

	for k := 4; k < top; k++ {

		next = make([][]int, 0)

		for _, old := range begin {
			temp := make([]int, len(old)+1)
			copy(temp, old)
			for i := 0; i < len(old); i++ {
				temp[i]++
			}
			reverse(temp[:len(temp)-1])

			for i := 1; i < len(old); i++ {
				mod := make([]int, len(temp))
				copy(mod, temp)
				reverse(mod[i:])

				next = append(next, mod)
			}

		}

		begin = next

	}

	sort.Sort(trains{next})
	fmt.Println("336/ Find the 2011th lexicographic maximix arrangement for eleven carriages")
	fmt.Println(letter(next[which-1]))
}

func reverse(list []int) {
	lth := len(list)
	for i := 0; i < lth/2; i++ {
		list[i], list[lth-i-1] = list[lth-i-1], list[i]
	}
}

func letter(list []int) (ans string) {
	for _, ch := range list {
		ans += string(ch + 65)
	}
	return
}
