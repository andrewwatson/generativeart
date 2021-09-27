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
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Black)
	c.SetLineWidth(1.0)
	c.SetLineColor(common.MediumAquamarine)
	c.SetAlpha(30)
	c.SetColorSchema(common.Plasma)
	c.SetIterations(4)
	c.FillBackground()
	c.Draw(arts.NewSilkSmoke(400, 20, 0.2, 2, 10, 30, false))
	c.ToPNG("silksmoke.png")
}
