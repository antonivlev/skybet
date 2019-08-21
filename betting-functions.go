/*
This file contains functions for playing different bets.
They must all take in a Roulette to play, and some parameters, e.g. bet amount,
bet number and return the winnings
*/
package main

import (
	"fmt"

	"github.com/antonivlev/skybet/roulette"
)

type SingleBetParams struct {
	BetNumber int
	BetAmount float64
}

// Returns customer's balance change; positive if win, negative if loss
// TODO: betNumber might not be in the roulette
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

// Rolls ball and gets the bet amount
// TODO: what if there is a colour mismatch, e.g. playColourBet(... "Red", ...) but roulette has "red"
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
