package roulette

import "testing"

func TestRoulette(t *testing.T) {
	r := Roulette{
		Numbers: []int{1, 2, 3, 4},
		Colours: []string{"r", "b", "r", "b"},
	}
	if false {
		t.Errorf("Broke %+v\n", r)
	}
}
