package arts

import (
	"math"
	"math/rand"

	"github.com/andrewwatson/generativeart"
	"github.com/andrewwatson/generativeart/common"
	"github.com/fogleman/gg"
)

type sinwavecircles struct {
	amplitude  float64
	spacing    float64
	noise      *common.PerlinNoise
	depth      int
	wavelength float64
}

// NewBlackHole returns a blackhole object.
func NewSinWaveCircles(amplitude int, wavelength float64, spacing int, depth int) *sinwavecircles {
	return &sinwavecircles{
		amplitude:  float64(amplitude),
		spacing:    float64(spacing),
		noise:      common.NewPerlinNoise(),
		depth:      depth,
		wavelength: wavelength,
	}
}

func (swc *sinwavecircles) Generative(c *generativeart.Canva) string {

	ctex := gg.NewContextForRGBA(c.Img())
	swc.recursionDraw(ctex, c, float64(c.Width()), swc.depth)
	return ""
}

func (swc *sinwavecircles) recursionDraw(ctex *gg.Context, c *generativeart.Canva, x float64, depth int) {
	if depth <= 0 {
		return
	}
	swc.draw(ctex, c, x)
	swc.recursionDraw(ctex, c, 1*x/4.0, depth-1)
	swc.recursionDraw(ctex, c, 2*x/4.0, depth-1)
	swc.recursionDraw(ctex, c, 3*x/4.0, depth-1)
}

// Generative draws a black hole image.
func (swc *sinwavecircles) draw(ctex *gg.Context, c *generativeart.Canva, x float64) {

	// noise := common.NewPerlinNoise()
	// kMax := common.RandomRangeFloat64(0.5, 1)
	ctex.SetLineWidth(0.4)
	ctex.SetColor(common.White)

	startingX := float64(c.Width()) * 0.05
	endingX := float64(c.Width()) * 0.95

	midpointY := float64(c.Height() / 2)

	noise := swc.noise.Noise3D(x*0.02+123.234, (1-x)*0.02, 345.4123)
	noise = math.Pow(noise, 0.5)
	amplitudeModulation := common.Remap(noise, 0.15, 0.85, 1.0, 20.0)

	radiusModulation := common.Remap(noise, 0.15, 0.85, 0.25, 1.25)

	for i := startingX; i < endingX; i += swc.spacing {

		var lw float64
		if rand.Float64() < 0.8 {
			lw = 1.2
		} else {
			lw = common.RandomRangeFloat64(1.0, common.RandomRangeFloat64(1, 3))
		}
		ctex.SetLineWidth(lw)

		radianX := gg.Radians(i * swc.wavelength)
		SinX := math.Sin(radianX) * swc.amplitude
		r := math.Pow(rand.Float64(), 2) * 100 * radiusModulation

		// fmt.Printf("X: %0.4f Rx: %0.4f Sin: %0.4f radius: %0.4f\n", i, radianX, SinX, r)

		cls := c.Opts().ColorSchema()[rand.Intn(len(c.Opts().ColorSchema()))]
		ctex.SetColor(cls)

		// inflationFactor := rand.Float64() * 25
		// yValue :=
		ctex.DrawCircle(i, midpointY+SinX*amplitudeModulation, r)
		ctex.Stroke()

	}

	ctex.ClearPath()

}
