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
	c := generativeart.NewCanva(1600, 1600)
	c.SetBackground(common.Azure)
	c.FillBackground()
	c.SetForeground(color.RGBA{113, 3, 0, 140})
	c.SetIterations(4000000)
	c.Draw(arts.NewSwirl(0.970, -1.899, 1.381, -1.506, 2.4, 2.4))
	c.ToPNG("swirl.png")
}
