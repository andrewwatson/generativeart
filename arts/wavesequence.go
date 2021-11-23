package arts

import (
	"image"
	"image/draw"
	"math"
	"math/rand"

	"github.com/andrewwatson/generativeart"
	"github.com/andrewwatson/generativeart/common"
	"github.com/fogleman/gg"
)

const (
	wavelength = 98.79
	// defaultAmplitude = 30.0
)

type WaveSequence struct {
	Options       WaveSequenceOptions
	MinimumRadius float64
	MaximumRadius float64
	startingY     float64
}

type WaveSequenceOptions struct {
	Frames int
}

func NewWaveSequence(options WaveSequenceOptions) *WaveSequence {

	return &WaveSequence{
		Options:       options,
		MinimumRadius: 5.0,
		MaximumRadius: 20.0,
		startingY:     -1.0,
	}
}
func (ws *WaveSequence) Generative(c *generativeart.Canva) string {

	// baseContext := gg.NewContextForRGBA(c.Img())

	ws.startingY = float64(c.Height()) / 2.0

	startingX := float64(c.Width()) * 0.05
	endingX := float64(c.Width()) * 0.95

	for i := startingX; i < endingX; i += defaultSpacing {

		last := c.GetLastFrame()
		bounds := last.Bounds()
		newFrame := image.NewRGBA(bounds)

		draw.Draw(newFrame, bounds, last, bounds.Min, draw.Src)

		ctex := gg.NewContextForRGBA(newFrame)

		var lw float64
		if rand.Float64() < 0.6 {
			lw = minLineWidth
		} else {
			lw = common.RandomRangeFloat64(minLineWidth, common.RandomRangeFloat64(minLineWidth, maxLineWidth))
		}
		ctex.SetLineWidth(lw)

		radianX := gg.Radians(i*wavelength) / 80
		SinX := math.Sin(radianX)
		amplitude := SinX * defaultAmplitude

		cls := c.Opts().ColorSchema()[rand.Intn(len(c.Opts().ColorSchema()))]
		ctex.SetColor(cls)
		radius := common.RandomRangeFloat64(ws.MinimumRadius, ws.MaximumRadius)
		ctex.DrawCircle(i, ws.startingY+(amplitude), radius)

		ctex.Stroke()

		c.AddFrame(newFrame)
	}
	return ""
}
