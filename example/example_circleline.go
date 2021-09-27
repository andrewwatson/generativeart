package main

import (
	"math/rand"
	"time"

	"github.com/andrewwatson/generativeart"
	"github.com/andrewwatson/generativeart/arts"
	"github.com/andrewwatson/generativeart/common"
)

func main() {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(600, 600)
	c.SetBackground(common.Tan)
	c.SetLineWidth(1.0)
	c.SetLineColor(common.LightPink)
	c.FillBackground()
	c.Draw(arts.NewCircleLine(0.02, 600, 1.5, 2, 2))
	c.ToPNG("circleline.png")
}
