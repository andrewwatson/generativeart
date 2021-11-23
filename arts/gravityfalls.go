package arts

import "github.com/andrewwatson/generativeart"

type gravityfalls struct {
	balls []ball
}

type ball struct {
	x      float64
	y      float64
	radius float64
}

type GravityFallsOptions struct {
	BallCount  int     // num balls
	Gravity    float64 // force of gravity in m/(s*s)
	Elasticity float64 // degree of elasticity when balls hit the floor.  0 = none, 1 = total
}

func (gf *gravityfalls) Generative(c *generativeart.Canva) {

}

func NewGravityFalls(options GravityFallsOptions) *gravityfalls {

	balls := make([]ball, options.BallCount)
	return &gravityfalls{
		balls: balls,
	}
}
