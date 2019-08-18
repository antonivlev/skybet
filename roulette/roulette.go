package roulette

type Roulette struct {
	Numbers []int
	Colours []string
}

func (r *Roulette) RollBall() (number int, colour string) {
	// this would be randomised
	ind := 0
	return r.Numbers[ind], r.Colours[ind]
}
