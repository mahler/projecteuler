package main

import (
	"fmt"
)

// In the United Kingdom the currency is made up of pound (£) and pence (p).
// There are eight coins in general circulation:
//
// 1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).
//
// It is possible to make £2 in the following way:
//
// 1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
//
// How many different ways can £2 be made using any number of coins?

func main() {
	numberOfWays := 0

	// loop for £2 (200p)
	for a := 0; a <= 1; a++ {
		baseSum200p := a * 200
		if baseSum200p >= 200 {
			if baseSum200p == 200 {
				numberOfWays++
			}
			break
		}

		// loop for £1 (100p)
		for b := 0; b <= 2; b++ {
			baseSum200plus100p := baseSum200p + b*100
			if baseSum200plus100p >= 200 {
				if baseSum200plus100p == 200 {
					fmt.Println("£2 = £2 *", a, "+ £1 *", b)
					numberOfWays++
				}
				// Sum above £2, finish loop as no further sum can be lower
				break
			}

			// loop for 50p
			for c := 0; c <= 4; c++ {
				baseSum50pAndLargerCoins := baseSum200plus100p + c*50
				if baseSum50pAndLargerCoins >= 200 {
					if baseSum50pAndLargerCoins == 200 {
						numberOfWays++
					}
					break
				}

				// loop for 20p
				for d := 0; d <= 10; d++ {
					baseSum20pAndLargerCoins := baseSum50pAndLargerCoins + d*20
					if baseSum20pAndLargerCoins >= 200 {
						if baseSum20pAndLargerCoins == 200 {
							numberOfWays++
						}
						break
					}

					// loop for 10p
					for e := 0; e <= 20; e++ {
						baseSum10pAndLargerCoins := baseSum20pAndLargerCoins + e*10
						if baseSum10pAndLargerCoins >= 200 {
							if baseSum10pAndLargerCoins == 200 {
								numberOfWays++
							}
							break
						}

						// loop for 5p
						for f := 0; f <= 40; f++ {
							baseSum5pAndLargerCoins := baseSum10pAndLargerCoins + f*5
							if baseSum5pAndLargerCoins >= 200 {
								if baseSum5pAndLargerCoins == 200 {
									numberOfWays++
								}
								break
							}

							// loop for 2p
							for g := 0; g <= 100; g++ {
								baseSum2pAndLargerCoins := baseSum5pAndLargerCoins + g*2
								if baseSum2pAndLargerCoins >= 200 {
									if baseSum2pAndLargerCoins == 200 {
										numberOfWays++
									}
									break
								}

								// loop for 1p
								for h := 0; h <= 200; h++ {
									coinSum := baseSum2pAndLargerCoins + h*1
									if coinSum >= 200 {
										if coinSum == 200 {
											numberOfWays++
										}
										break
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("031/ How many different ways can £2 be made using any number of coins?")
	fmt.Println(numberOfWays)
}
