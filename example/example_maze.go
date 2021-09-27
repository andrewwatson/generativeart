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
	c.SetBackground(common.Azure)
	c.SetLineWidth(3)
	c.SetLineColor(common.Orange)
	c.FillBackground()
	c.Draw(arts.NewMaze(20))
	c.ToPNG("maze.png")
}
