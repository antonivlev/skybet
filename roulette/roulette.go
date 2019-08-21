/*
Provides a roulette and a method to play on it (RollBall)
*/
package roulette

import (
	"math/rand"
	"time"
)

/*
Contains information for a simple two colour roulette.

TODO, constraints:
1) Same number cannot be both red and black (make setters for the numbers)
2) Can't be empty?
*/
type Roulette struct {
	BlackNumbers []int
	RedNumbers   []int
}

func (r *Roulette) RollBall() (number int, colour string) {
	allNumbers := append(r.BlackNumbers, r.RedNumbers...)

	// Use current time to enable randomness every time
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Pick random number from roulette
	randomIndex := randSource.Intn(len(allNumbers))
	chosenNumber := allNumbers[randomIndex]

	// Assign colour based on where the chosen number is
	colour = "black"
	if contains(r.RedNumbers, chosenNumber) {
		colour = "red"
	}

	return chosenNumber, colour
}

// Utlity func: array containment
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
