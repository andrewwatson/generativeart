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
	colors := []color.RGBA{
		{0x11, 0x60, 0xC6, 0xFF},
		{0xFD, 0xD9, 0x00, 0xFF},
		{0xF5, 0xB4, 0xF8, 0xFF},
		{0xEF, 0x13, 0x55, 0xFF},
		{0xF4, 0x9F, 0x0A, 0xFF},
	}
	c := generativeart.NewCanva(800, 800)
	c.SetBackground(common.White)
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewColorCircle2(30))
	c.ToPNG("colorcircle2.png")
}
