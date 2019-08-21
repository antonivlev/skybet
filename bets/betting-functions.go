/*
This package contains functions for playing different bets.
They must all be of type BettingFunc:

type BettingFunc func(*roulette.Roulette, BetArgs) (float64, int, string)

TODO: for testing this package, would be nice to take each function in the package and:
1) Make sure it's of the BettingFunc type
2) Try running it with several different roulettes and BetArgs

Pretty sure a script could be written to do this automatically, possibly using reflection
*/
package bets

import (
	"github.com/antonivlev/skybet/roulette"
)

/*
Generic betting func type. The type is needed to avoid having to parse
parameters individually, as different betting funcs will have different
signatures.
*/
type BettingFunc func(*roulette.Roulette, BetArgs) (float64, int, string)

/*
Contains all possible arguments for all bets. Whatever the frontend passes through will be
decoded into this struct. Unpack in your betting function what you need.

Example name mapping:
	"iseven", "isEven", "IsEven" in URL would all map to an IsEven field in the struct

NOTE: not ideal, groups all possible parameters for no other reason than to avoid individual
parsing.
*/
type BetArgs struct {
	// all bets have an amount of money bet on
	Money float64
	// used for single number bet
	Number int
	// used in a colour bet
	Colour string
}

// Plays bet on a single number, needs Money and Number
// TODO: betNumber might not be in the roulette
func PlayBetOnSingleNumber(r *roulette.Roulette, args BetArgs) (win float64, outNum int, outCol string) {
	// unpack the args you need
	betAmount, betNumber := args.Money, args.Number

	// roll the ball (don't need colour in this bet type, used only to pass back to user)
	outNum, outCol = r.RollBall()

	// winning conditions
	if outNum == betNumber {
		// winning calculation:
		// multiplier here based on roulette win chance and what casino decides
		allNumbers := append(r.BlackNumbers, r.RedNumbers...)
		win = betAmount * float64(len(allNumbers)-2)
	} else {
		win = 0.0
	}

	return win, outNum, outCol
}

// Plays bet on a colour, needs Money and Colour
// TODO: what if there is a colour mismatch, e.g. args.Colour = "redd", but roulette has "red"
func PlayColourBet(r *roulette.Roulette, args BetArgs) (win float64, outNum int, outCol string) {
	betAmount, betColour := args.Money, args.Colour
	outNum, outCol = r.RollBall()

	if betColour == outCol {
		win = betAmount * 2
	} else {
		win = 0.0
	}

	return win, outNum, outCol
}
