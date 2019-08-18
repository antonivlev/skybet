package main

import (
	"fmt"

	"github.com/antonivlev/skybet/roulette"
)

func main() {
	r := roulette.Roulette{
		Numbers: []int{1, 2, 3, 4},
		Colours: []string{"red", "black", "red", "black"},
	}
	fmt.Printf("roulette: \n%+v\n\n", r)

	win := placeBetOnSingleNumber(&r, 3, 13.50)
	fmt.Printf("I won %4.2f\n", win)
}

// Returns customer's balance change; positive if win, negative if loss
func placeBetOnSingleNumber(r *roulette.Roulette, betNumber int, amount float64) (win float64) {
	// roll the ball (don't need colour in this bet type, used only to print)
	n, c := r.RollBall()
	fmt.Printf("rolled, got: %d %s\n", n, c)

	// winning conditions
	if n == betNumber {
		// winning calculation:
		// multiplier here based on roulette and what casino decides
		// standard: 35, if there are 37 betting options
		win = amount * 2
	} else {
		win = -amount
	}
	return win
}
