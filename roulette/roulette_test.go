package roulette

import "testing"

func TestRollOutput(t *testing.T) {
	r := Roulette{
		BlackNumbers: []int{1, 2, 3, 4},
		RedNumbers:   []int{5, 6, 7, 8},
	}

	// Check roll output is within roulette
	numbers := append(r.BlackNumbers, r.RedNumbers...)
	n, c := r.RollBall()
	if !contains(numbers, n) {
		t.Errorf("Rolling gave %d, but it's not in %+v", n, r)
	} else if c != "red" && c != "black" {
		t.Errorf("The colour '%s' is neither red nor black", c)
	}
}
