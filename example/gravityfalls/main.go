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

const (
	minRadius = 10
	maxRadius = 30
)

func main() {

	colors := []color.RGBA{
		{0xF9, 0xC8, 0x0E, 0xFF},
		{0xF8, 0x66, 0x24, 0xFF},
		{0xEA, 0x35, 0x46, 0xFF},
		{0x66, 0x2E, 0x9B, 0xFF},
		{0x43, 0xBC, 0xCD, 0xFF},
	}

	size := 1000
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(size, size)

	c.SetColorSchema(colors)

	c.SetBackground(common.Black)
	c.FillBackground()

	waveOptions := arts.WaveSequenceOptions{
		Frames: 10,
	}

	c.Draw(arts.NewWaveSequence(waveOptions))
	fileID := "a8c5e8d7-f100-4d8a-871a-bc5c1a8cf037"
	fileName := fmt.Sprintf("%s.gif", fileID)
	err := c.ToAnimatedGIF(fileName, 2, -1)
	if err != nil {
		panic(err)
	}
}