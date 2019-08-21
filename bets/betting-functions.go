/*
This file contains functions for playing different bets.
They must all take in a Roulette to play, and some parameters, e.g. bet amount,
bet number and return the winnings
*/
package bets

import (
	"fmt"

	"github.com/antonivlev/skybet/roulette"
)

type BettingFunc func(*roulette.Roulette, BetArgs) (float64, string)

// Contains all possible arguments for all bets. Unpack in your betting function as needed.
type BetArgs struct {
	// all bets have an amount of money bet on
	Money float64
	// used for single number bet
	Number int
	// used in a colour bet
	Colour string
}

// Returns customer's balance change; positive if win, negative if loss
// TODO: betNumber might not be in the roulette
func PlayBetOnSingleNumber(r *roulette.Roulette, args BetArgs) (win float64, msg string) {
	// unpack the args you need
	betAmount, betNumber := args.Money, args.Number

	// roll the ball (don't need colour in this bet type, used only for msg)
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

	msg = fmt.Sprintf("bet %4.2f on %d:\ngot: %d %s\n", betAmount, betNumber, n, c)
	return win, msg
}

// Rolls ball and gets the bet amount
// TODO: what if there is a colour mismatch, e.g. playColourBet(... "Red", ...) but roulette has "red"
func PlayColourBet(r *roulette.Roulette, args BetArgs) (win float64, msg string) {
	betAmount, colour := args.Money, args.Colour
	n, c := r.RollBall()

	if colour == c {
		win = betAmount * 2
	} else {
		win = -betAmount
	}

	msg = fmt.Sprintf("bet %4.2f on %s:\ngot: %d %s\n", betAmount, colour, n, c)
	return win, msg
}
