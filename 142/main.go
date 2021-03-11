package main

import (
	"fmt"
	"math"
)

const limit2 = 5000
const count = 200000

func main() {
	C2Table := make(map[int64]map[int64]bool)

	for a := int64(1); a < limit2; a++ {
		for b := a; b < limit2; b++ {
			c2 := a*a + b*b
			if IsSquare(c2) {
				if enter, ok := C2Table[c2]; ok {
					enter[a*a] = true
					enter[b*b] = true
				} else {
					C2Table[c2] = make(map[int64]bool, 0)
					C2Table[c2][a*a] = true
					C2Table[c2][b*b] = true
				}

			}

		}
	}

	//fmt.Println("Table Built")

	min := int64(3*count + 3*limit2*limit2)

	for x := int64(1); x < count; x++ {
		for c2, poss := range C2Table {
			z := x + c2

			if x+x+z < min && IsSquare(x+z) { //x+x+z < x+y+z < min
				for a2 := range poss {
					y := x + a2
					if x+y+z < min && IsSquare(x+y) && IsSquare(z+y) {

						min = x + y + z
						//fmt.Println(x, "+", y, "+", z, "=", x+y+z)
					}
				}
			}
		}
	}
	fmt.Println("142/")
	fmt.Println(min)
}

func IsSquare(x int64) bool {
	return int64(math.Sqrt(float64(x)))*int64(math.Sqrt(float64(x))) == x
}
