package arts

import (
	"math"
	"math/rand"

	"github.com/andrewwatson/generativeart"
	"github.com/andrewwatson/generativeart/common"
	"github.com/fogleman/gg"
)

type noiseLine struct {
	n       int
	elipses bool
}

// NewNoiseLine returns a noiseLine object.
func NewNoiseLine(n int, elipses bool) *noiseLine {
	return &noiseLine{
		n:       n,
		elipses: elipses,
	}
}

// Generative draws a noise line image.
func (nl *noiseLine) Generative(c *generativeart.Canva) string {
	ctex := gg.NewContextForRGBA(c.Img())
	noise := common.NewPerlinNoise()

	ctex.SetColor(common.Black)
	if nl.elipses {
		for i := 0; i < 80; i++ {
			x := rand.Float64() * float64(c.Width())
			y := rand.Float64() * float64(c.Height())

			s := rand.Float64() * float64(c.Width()) / 8
			ctex.SetLineWidth(0.5)
			ctex.DrawEllipse(x, y, s, s)
			ctex.Stroke()
		}
	}

	t := rand.Float64() * 10
	for i := 0; i < nl.n; i++ {
		x := common.RandomRangeFloat64(-0.5, 1.5) * float64(c.Width())
		y := common.RandomRangeFloat64(-0.5, 1.5) * float64(c.Height())
		cl := c.Opts().ColorSchema()[rand.Intn(len(c.Opts().ColorSchema()))]
		cl.A = uint8(c.Opts().Alpha())

		l := 400
		for j := 0; j < l; j++ {
			var ns = 0.0005
			w := math.Sin(math.Pi*float64(j)/float64(l-1)) * 5
			theta := noise.Noise3D(x*ns, y*ns, t) * 100
			ctex.SetColor(cl)
			ctex.DrawCircle(x, y, w)
			ctex.Fill()
			x += math.Cos(theta)
			y += math.Sin(theta)
			t += 0.0000003
		}
	}
	return ""
}
