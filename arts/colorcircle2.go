package arts

import (
	"math"
	"math/rand"

	"github.com/andrewwatson/generativeart"
	"github.com/andrewwatson/generativeart/common"
	"github.com/fogleman/gg"
)

type colorCircle2 struct {
	circleNum int
}

// NewColorCircle2 returns a colorCircle2 object.
func NewColorCircle2(circleNum int) *colorCircle2 {
	return &colorCircle2{
		circleNum: circleNum,
	}
}

// Generative draws a color circle image.
func (cc *colorCircle2) Generative(c *generativeart.Canva) string {
	ctex := gg.NewContextForRGBA(c.Img())

	for i := 0; i < cc.circleNum; i++ {
		x := common.RandomRangeFloat64(0, float64(c.Width()))
		y := common.RandomRangeFloat64(0, float64(c.Height()))

		r1 := common.RandomRangeFloat64(50.0, float64(c.Width())/4)
		r2 := common.RandomRangeFloat64(10.0, float64(c.Width())/3)

		cc.circle(ctex, c, x, y, r1, r2)
		if rand.Float64() < 0.3 {
			col := c.Opts().ColorSchema()[rand.Intn(len(c.Opts().ColorSchema()))]
			ctex.SetColor(col)
			ctex.DrawCircle(x, y, r1/2.0)
			ctex.Fill()
		}
	}
	return ""
}

func (cc *colorCircle2) circle(ctex *gg.Context, c *generativeart.Canva, x, y, d, dx float64) {
	col := c.Opts().ColorSchema()[rand.Intn(len(c.Opts().ColorSchema()))]

	for j := 0.0; j < dx; j += 1.0 {
		dd := d + j*2.0
		alpha := int((dx / j) * 255.0)
		if alpha > 255 {
			alpha = 255
		}
		col.A = uint8(alpha)

		ctex.SetColor(col)
		for i := 0; i < 150; i++ {
			theta := common.RandomRangeFloat64(0, math.Pi*2)
			xx := x + dd*0.5*math.Cos(theta)
			yy := y + dd*0.5*math.Sin(theta)
			ctex.DrawPoint(xx, yy, 0.51)
			ctex.Fill()
		}
	}
}
