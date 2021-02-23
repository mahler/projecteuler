package main

import (
	"fmt"
)

// A bag contains one red disc and one blue disc. In a game of chance a player takes a disc at random and its colour is noted.
// After each turn the disc is returned to the bag, an extra red disc is added, and another disc is taken at random.
//
// The player pays £1 to play and wins if they have taken more blue discs than red discs at the end of the game.
//
// If the game is played for four turns, the probability of a player winning is exactly 11/120, and so the maximum prize fund
// the banker should allocate for winning in this game would be £10 before they would expect to incur a loss. Note that any payout
// will be a whole number of pounds and also includes the original £1 paid to play the game, so in the example given the player
// actually wins £9.
//
// Find the maximum prize fund that should be allocated to a single game in which fifteen turns are played.
func main() {
	fmt.Println("121/ Find the maximum prize fund that should be allocated to a single game in which fifteen turns are played.")
	fmt.Println(int(1 / prob(0, 0, 15)))
}

func prob(won, drawn, rounds int) float64 {

	if drawn == rounds {
		if 2*won > drawn {
			return 1
		} else {
			return 0
		}
	}

	return prob(won+1, drawn+1, rounds)*1/float64(drawn+2) + prob(won, drawn+1, rounds)*float64(drawn+1)/float64(drawn+2)

}
