package main

import (
	"fmt"
)

// Peter has nine four-sided (pyramidal) dice, each with faces numbered 1, 2, 3, 4.
// Colin has six six-sided (cubic) dice, each with faces numbered 1, 2, 3, 4, 5, 6.
//
// Peter and Colin roll their dice and compare totals: the highest total wins.
// The result is a draw if the totals are equal.
//
// What is the probability that Pyramidal Pete beats Cubic Colin?
// Give your answer rounded to seven decimal places in the form 0.abcdefg

func main() {
	fmt.Println("Patience...")
	var wins, tries int64
	for p1 := 1; p1 < 5; p1++ {
		for p2 := 1; p2 < 5; p2++ {
			for p3 := 1; p3 < 5; p3++ {
				for p4 := 1; p4 < 5; p4++ {
					for p5 := 1; p5 < 5; p5++ {
						for p6 := 1; p6 < 5; p6++ {
							for p7 := 1; p7 < 5; p7++ {
								for p8 := 1; p8 < 5; p8++ {
									for p9 := 1; p9 < 5; p9++ {
										for c1 := 1; c1 < 7; c1++ {
											for c2 := 1; c2 < 7; c2++ {
												for c3 := 1; c3 < 7; c3++ {
													for c4 := 1; c4 < 7; c4++ {
														for c5 := 1; c5 < 7; c5++ {
															for c6 := 1; c6 < 7; c6++ {

																tries++
																if p1+p2+p3+p4+p5+p6+p7+p8+p9 > c1+c2+c3+c4+c5+c6 {
																	wins++
																}

															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}

				}
			}
		}
	}

	result := float64(wins) / float64(tries)
	strResult := fmt.Sprintf("%.7f", result)
	fmt.Println("205/ What is the probability that Pyramidal Pete beats Cubic Colin?")
	fmt.Println(strResult)
}
