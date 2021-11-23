package arts

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/andrewwatson/generativeart"
	"github.com/andrewwatson/generativeart/common"
	"github.com/fogleman/gg"
)

type circleLoop2 struct {
	depth int
	noise *common.PerlinNoise
}

func NewCircleLoop2(depth int) *circleLoop2 {
	return &circleLoop2{
		depth: depth,
		noise: common.NewPerlinNoise(),
	}
}

// Generative draws a circle composed by many colored circles.
func (cl *circleLoop2) Generative(c *generativeart.Canva) {
	ctex := gg.NewContextForRGBA(c.Img())
	ctex.Translate(float64(c.Width())/2, float64(c.Height())/2)

	numCircles := cl.recursionDraw(ctex, c, float64(c.Width()), cl.depth)
	fmt.Printf("Total Circles: %d\n", numCircles)
}

func (cl *circleLoop2) recursionDraw(ctex *gg.Context, c *generativeart.Canva, x float64, depth int) int {
	if depth <= 0 {
		return 0
	}

	circles := 0
	circles += cl.draw(ctex, c, x)
	circles += cl.recursionDraw(ctex, c, 1*x/4.0, depth-1)
	circles += cl.recursionDraw(ctex, c, 2*x/4.0, depth-1)
	circles += cl.recursionDraw(ctex, c, 3*x/4.0, depth-1)

	return circles
}

func (cl *circleLoop2) draw(ctex *gg.Context, c *generativeart.Canva, x float64) int {
	var lw float64
	odds := rand.Float64()
	if odds < 0.8 {
		lw = 2
	} else if odds > 0.95 {
		lw = 5
	} else {
		lw = common.RandomRangeFloat64(2.0, common.RandomRangeFloat64(1, 6))
	}
	ctex.SetLineWidth(lw)

	noise := cl.noise.Noise3D(x*0.02+123.234, (1-x)*0.02, 345.4123)
	noise = math.Pow(noise, 0.5)
	a2 := common.Remap(noise, 0.15, 0.85, 0.1, 0.6)

	px := math.Pow(x/float64(c.Height()), a2) * float64(c.Height())
	py := math.Pow(1-x/float64(c.Height()), a2)*float64(c.Height()) -
		common.RandomRangeFloat64(0,
			common.RandomRangeFloat64(float64(c.Height())*0.18,
				common.RandomRangeFloat64(float64(c.Height())*0.18,
					float64(c.Height())*0.7,
				),
			),
		)

	cls := c.Opts().ColorSchema()[rand.Intn(len(c.Opts().ColorSchema()))]
	ctex.SetColor(cls)
	nCircles := common.RandomRangeInt(1, 6)
	if rand.Float64() < 0.03 {
		nCircles = common.RandomRangeInt(8, 10)
	}

	r := math.Pow(rand.Float64(), 2) * 50

	var flag bool
	if rand.Float64() < 0.7 {
		flag = true
	}

	for i := 0; i < nCircles; i++ {
		if flag {
			ctex.DrawCircle(px*0.39, py*0.39, rand.Float64()*float64(i)*r/float64(nCircles))
		} else {
			ctex.DrawCircle(px*0.39, py*0.39, float64(i)*r/float64(nCircles))
		}
		ctex.Stroke()
	}
	ctex.Rotate(x / float64(c.Height()) * 0.2)

	return nCircles
}
