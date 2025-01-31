package arts

import (
	"math"
	"math/rand"

	"github.com/andrewwatson/generativeart"
	"github.com/andrewwatson/generativeart/common"
	"github.com/fogleman/gg"
)

type point struct {
	x, y float64
}

type circleLine struct {
	step         float64
	lineNum      int
	radius       float64
	xaixs, yaixs float64
}

// NewCircleLine returns a circleLine object.
func NewCircleLine(step float64, lineNum int, radius, xaixs, yaixs float64) *circleLine {
	return &circleLine{
		step:    step,
		lineNum: lineNum,
		radius:  radius,
		xaixs:   xaixs,
		yaixs:   yaixs,
	}
}

// Generative draws a cirle line image.
func (cl *circleLine) Generative(c *generativeart.Canva) string {
	ctex := gg.NewContextForRGBA(c.Img())
	ctex.SetLineWidth(c.Opts().LineWidth())
	ctex.SetColor(c.Opts().LineColor())
	var points []point
	for theta := -math.Pi; theta <= math.Pi; theta += cl.step {
		x := cl.radius * math.Cos(theta)
		y := cl.radius * math.Sin(theta)
		xi, yi := common.ConvertCartesianToPixel(x, y, cl.xaixs, cl.yaixs, c.Width(), c.Height())
		points = append(points, point{
			x: float64(xi),
			y: float64(yi),
		})
	}

	for i := 0; i < cl.lineNum; i++ {
		p1 := points[rand.Intn(len(points))]
		p2 := points[rand.Intn(len(points))]

		ctex.MoveTo(p1.x, p1.y)
		ctex.LineTo(p2.x, p2.y)
		ctex.Stroke()
	}

	return ""
}
