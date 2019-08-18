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

	win1 := playBetOnSingleNumber(&r, 3, 13.50)
	fmt.Printf("  result: %4.2f\n", win1)
	win2 := playBetOnSingleNumber(&r, 3, 13.50)
	fmt.Printf("  result: %4.2f\n", win2)
	win3 := playBetOnSingleNumber(&r, 3, 13.50)
	fmt.Printf("  result: %4.2f\n", win3)
}

// Returns customer's balance change; positive if win, negative if loss
func playBetOnSingleNumber(r *roulette.Roulette, betNumber int, betAmount float64) (win float64) {
	// roll the ball (don't need colour in this bet type, used only to print)
	n, c := r.RollBall()
	fmt.Printf("bet %4.2f on %d, got: %d %s,", betAmount, betNumber, n, c)

	// winning conditions
	if n == betNumber {
		// winning calculation:
		// multiplier here based on roulette win chance and what casino decides
		allNumbers := append(r.BlackNumbers, r.RedNumbers...)
		win = betAmount * float64(len(allNumbers)-2)
	} else {
		win = -betAmount
	}
	return win
}
