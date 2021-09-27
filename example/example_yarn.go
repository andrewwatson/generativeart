package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/andrewwatson/generativeart"
	"github.com/andrewwatson/generativeart/arts"
	"github.com/andrewwatson/generativeart/common"
)

func main() {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.Orange)
	c.FillBackground()
	c.SetLineWidth(0.3)
	c.SetLineColor(color.RGBA{A: 60})
	c.Draw(arts.NewYarn(2000))
	c.ToPNG("yarn.png")
}
