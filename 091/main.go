package main

import (
	"fmt"
)

// See https://projecteuler.net/problem=91 for problem description

func main() {
	max := 50
	total := 0

	for x1 := 0; x1 <= max; x1++ {
		for y1 := 0; y1 <= max; y1++ {
			for x2 := 0; x2 <= max; x2++ {
				for y2 := 0; y2 <= max; y2++ {

					//if points are multiples there is no valid triangle
					if x1*y2 != x2*y1 {
						if ((x1*x2)+(y1*y2)) == 0 || (x1*(x2-x1)+y1*(y2-y1)) == 0 || (x2*(x2-x1)+y2*(y2-y1)) == 0 {
							total++
						}
					}

				}
			}
		}
	}

	fmt.Println("091/how many right triangles can be formed?")
	fmt.Println(total / 2)
}
