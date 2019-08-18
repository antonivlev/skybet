package main

import (
	"fmt"

	"github.com/antonivlev/skybet/roulette"
)

func main() {
	r := roulette.Roulette{
		RedNumbers:   []int{1, 2, 3, 4},
		BlackNumbers: []int{5, 6, 7, 8},
	}
	fmt.Printf("roulette: \n%+v\n\n", r)

	playBetOnSingleNumber(&r, 3, 13.50)
	playBetOnSingleNumber(&r, 3, 13.50)
	playBetOnSingleNumber(&r, 3, 13.50)

	playColourBet(&r, "red", 13.50)
	playColourBet(&r, "red", 13.50)
	playColourBet(&r, "black", 13.50)
}

// Returns customer's balance change; positive if win, negative if loss
func playBetOnSingleNumber(r *roulette.Roulette, betNumber int, betAmount float64) (win float64) {
	// roll the ball (don't need colour in this bet type, used only to print)
	n, c := r.RollBall()

	// winning conditions
	if n == betNumber {
		// winning calculation:
		// multiplier here based on roulette win chance and what casino decides
		allNumbers := append(r.BlackNumbers, r.RedNumbers...)
		win = betAmount * float64(len(allNumbers)-2)
	} else {
		win = -betAmount
	}

	fmt.Printf("bet %4.2f on %d, got: %d %s,  result: %4.2f\n", betAmount, betNumber, n, c, win)
	return win
}

func playColourBet(r *roulette.Roulette, colour string, betAmount float64) (win float64) {
	n, c := r.RollBall()

	if colour == c {
		win = betAmount * 2
	} else {
		win = -betAmount
	}

	fmt.Printf("bet %4.2f on %s, got: %d %s,  result: %4.2f\n", betAmount, colour, n, c, win)
	return win
}
