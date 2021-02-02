package main

import "fmt"

// Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:
//
// 21 22 23 24 25
// 20  7  8  9 10
// 19  6  1  2 11
// 18  5  4  3 12
// 17 16 15 14 13
//
// It can be verified that the sum of the numbers on the diagonals is 101.
//
// What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
func main() {
	size := 1001

	spiral := make([][]int, size)
	for i, _ := range spiral {
		spiral[i] = make([]int, size)
	}

	// ----
	toggle := 0
	moves := 1
	// 0 = right, 1 = down, 2 = left, 3 = up
	direction := 0
	v, i, j := 1, size/2, size/2

	for v < size*size {
		for n := 0; n < moves; n++ {
			spiral[i][j] = v
			switch direction {
			case 0:
				j++
			case 1:
				i++
			case 2:
				j--
			case 3:
				i--
			}
			v++
		}
		direction = (direction + 1) % 4
		if toggle > 0 {
			moves++
			toggle = 0
		} else {
			toggle++
		}
	}

	fmt.Println("028/ // What is the sum of the diagonals in spiral formed in the same way?")
	fmt.Println(diagonal_sum(spiral))
}

func ceildiv(n, d int) int {
	if n%d == 0 {
		return n / d
	}
	return (n / d) + 1
}

func diagonal_sum(spiral [][]int) int {
	size := len(spiral)
	sum := 0
	for i, j := 0, 0; i < size; {
		sum += spiral[i][j]
		i++
		j++
	}
	for i, j := 0, size-1; i < size; {
		sum += spiral[i][j]
		i++
		j--
	}
	// adjust for counting the middle twice
	sum -= 1
	return sum
}
