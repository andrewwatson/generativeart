package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/andrewwatson/generativeart"
	"github.com/andrewwatson/generativeart/arts"
	"github.com/andrewwatson/generativeart/common"
)

func main() {

	colors := []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}

	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(2000, 2000)

	c.SetColorSchema(colors)

	c.SetBackground(common.Black)
	c.FillBackground()
	c.SetLineWidth(4.0)
	c.SetLineColor(common.White)

	amplitude := 12
	spacing := 2
	depth := 2

	wavelength := 1.0

	c.Draw(arts.NewSinWaveCircles(amplitude, wavelength, spacing, depth))
	fileID := "a8c5e8d7-f100-4d8a-871a-bc5c1a8cf037"
	fileName := fmt.Sprintf("%s.png", fileID)
	c.ToPNG(fileName)
}
