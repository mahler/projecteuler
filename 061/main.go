package main

import (
	"fmt"
	"strconv"
)

func main() {
	lists := make([][]int, 6)

	for permutation := 0; permutation < 720 && len(lists[0]) == 0; permutation++ {
		order := Permutation(permutation, []int{0, 1, 2, 3, 4, 5})
		quantity := len(order)

		i := 1
		for ; triangle(i) < 1000; i++ {
		}
		for ; triangle(i) < 10000; i++ {
			lists[order[0]] = append(lists[order[0]], triangle(i))
		}

		i = 1
		for ; square(i) < 1000; i++ {
		}
		for ; square(i) < 10000; i++ {
			lists[order[1]] = append(lists[order[1]], square(i))
		}

		i = 1
		for ; pentagon(i) < 1000; i++ {
		}
		for ; pentagon(i) < 10000; i++ {
			lists[order[2]] = append(lists[order[2]], pentagon(i))
		}

		i = 1
		for ; hexagon(i) < 1000; i++ {
		}
		for ; hexagon(i) < 10000; i++ {
			lists[order[3]] = append(lists[order[3]], hexagon(i))
		}

		i = 1
		for ; heptagon(i) < 1000; i++ {
		}
		for ; heptagon(i) < 10000; i++ {
			lists[order[4]] = append(lists[order[4]], heptagon(i))
		}

		i = 1
		for ; octagon(i) < 1000; i++ {
		}
		for ; octagon(i) < 10000; i++ {
			lists[order[5]] = append(lists[order[5]], octagon(i))
		}

		for place := 1; place < 200 && len(lists[0]) != 0; place++ {

			for i := 0; i < len(lists[place%quantity]); i++ {
				matches := false
				for j := 0; j < len(lists[(place+1)%quantity]) && !matches; j++ {

					matches = links(lists[place%quantity][i], lists[(place+1)%quantity][j])

				}

				if !matches {
					lists[place%quantity] = append(lists[place%quantity][:i], lists[place%quantity][i+1:]...)
				}

			}

			for i := 0; i < len(lists[place%quantity]); i++ {
				matches := false
				for j := 0; j < len(lists[(place-1)%quantity]) && !matches; j++ {
					matches = links(lists[(place-1)%quantity][j], lists[place%quantity][i])
				}

				if !matches {
					lists[place%quantity] = append(lists[place%quantity][:i], lists[place%quantity][i+1:]...)
				}
			}
		}

		if len(lists[0]) != 0 {
			sum := 0
			for i := 0; i < quantity; i++ {
				sum += lists[i][0]
			}
			fmt.Println("061/ Find the sum of the only ordered set of six cyclic 4-digit numbers")
			fmt.Println("for which each polygonal type:")
			fmt.Println(sum)
		}
	}
}

func triangle(n int) int {
	return n * (n + 1) / 2
}

func square(n int) int {
	return n * n
}

func pentagon(n int) int {
	return n * ((3 * n) - 1) / 2
}

func hexagon(n int) int {
	return n * ((2 * n) - 1)
}

func heptagon(n int) int {
	return n * ((5 * n) - 3) / 2
}

func octagon(n int) int {
	return n * ((3 * n) - 2)
}

func links(a int, b int) bool {
	A := strconv.Itoa(a)
	B := strconv.Itoa(b)
	return A[2:] == B[:2]

}

//Permutation returns the nth permutation of a slice of integer
//values. Undefined behavior for n > (len(list) factorial).
func Permutation(n int, list []int) []int {
	if len(list) == 1 {
		return list
	}

	k := n % len(list)

	first := []int{list[k]}
	next := make([]int, len(list)-1)

	copy(next, append(list[:k], list[k+1:]...))

	return append(first, Permutation(n/len(list), next)...)
}
