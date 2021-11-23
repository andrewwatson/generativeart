package generativeart

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/andrewwatson/generativeart/common"
	"github.com/andybons/gogif"
)

type Engine interface {
	Generative(c *Canva)
}

type Canva struct {
	height, width int
	img           *image.RGBA
	opts          Options
	frames        []*image.RGBA
}

func (c *Canva) Opts() Options {
	return c.opts
}

func (c *Canva) Img() *image.RGBA {
	return c.img
}

func (c *Canva) Width() int {
	return c.width
}

func (c *Canva) Height() int {
	return c.height
}

type Options struct {
	background  color.RGBA
	foreground  color.RGBA
	lineColor   color.RGBA
	lineWidth   float64
	colorSchema []color.RGBA
	nIters      int
	alpha       int
}

func (o Options) Alpha() int {
	return o.alpha
}

func (o Options) NIters() int {
	return o.nIters
}

func (o Options) ColorSchema() []color.RGBA {
	return o.colorSchema
}

func (o Options) LineWidth() float64 {
	return o.lineWidth
}

func (o Options) LineColor() color.RGBA {
	return o.lineColor
}

func (o Options) Foreground() color.RGBA {
	return o.foreground
}

func (o Options) Background() color.RGBA {
	return o.background
}

// NewCanva returns a Canva.
func NewCanva(w, h int) *Canva {
	return &Canva{
		height: h,
		width:  w,
		img:    image.NewRGBA(image.Rect(0, 0, w, h)),
		// Set some defaults value
		opts: Options{
			background:  common.Azure,
			foreground:  common.MistyRose,
			lineColor:   common.Tomato,
			lineWidth:   3,
			colorSchema: common.Youthful,
			nIters:      20,
			alpha:       255,
		},
		frames: make([]*image.RGBA, 0),
	}
}

func (c *Canva) SetOptions(opts Options) {
	c.opts = opts
}

func (c *Canva) SetBackground(rgba color.RGBA) {
	c.opts.background = rgba
}

func (c *Canva) SetForeground(rgba color.RGBA) {
	c.opts.foreground = rgba
}

func (c *Canva) SetColorSchema(rgbas []color.RGBA) {
	c.opts.colorSchema = rgbas
}

func (c *Canva) SetLineColor(rgba color.RGBA) {
	c.opts.lineColor = rgba
}

func (c *Canva) SetLineWidth(lw float64) {
	c.opts.lineWidth = lw
}

func (c *Canva) SetIterations(nIters int) {
	c.opts.nIters = nIters
}

func (c *Canva) SetAlpha(alpha int) {
	c.opts.alpha = alpha
}

func (c *Canva) Draw(e Engine) {
	e.Generative(c)
}

func (c *Canva) DrawTimelapse(e Engine) {

}

// FillBackground fills the background of the Canva.
func (c *Canva) FillBackground() {
	draw.Draw(c.Img(), c.Img().Bounds(), &image.Uniform{c.Opts().Background()}, image.ZP, draw.Src)
}

func (c *Canva) AddFrame(img *image.RGBA) {

	c.frames = append(c.frames, img)
}

func (c *Canva) GetLastFrame() *image.RGBA {
	if len(c.frames) > 0 {
		return c.frames[len(c.frames)-1]
	}

	return c.Img()
}

func (c *Canva) ToAnimatedGIF(fpath string, frameRate, loopCount int) error {

	outGif := &gif.GIF{}

	quantizer := gogif.MedianCutQuantizer{NumColor: 64}

	for _, simage := range c.frames {
		bounds := simage.Bounds()
		palettedImage := image.NewPaletted(bounds, nil)

		quantizer.Quantize(palettedImage, bounds, simage, image.ZP)

		// Add new frame to animated GIF
		outGif.Image = append(outGif.Image, palettedImage)
		outGif.Delay = append(outGif.Delay, frameRate)
	}

	f, err := os.Create(fpath)
	if err != nil {
		return err
	}
	outGif.LoopCount = loopCount
	gif.EncodeAll(f, outGif)

	return nil
}

func (c *Canva) ToGIF(fpath string) error {

	f, err := os.Create(fpath)
	if err != nil {
		return err
	}
	if err = gif.Encode(f, c.Img(), &gif.Options{
		NumColors: 256,
	}); err != nil {
		f.Close()
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}

// ToPng saves the image to local with PNG format.
func (c *Canva) ToPNG(fpath string) error {
	f, err := os.Create(fpath)
	if err != nil {
		return err
	}
	if err = png.Encode(f, c.Img()); err != nil {
		f.Close()
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}

// ToJpeg saves the image to local with Jpeg format.
func (c *Canva) ToJPEG(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := jpeg.Encode(f, c.Img(), nil); err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

// ToBytes returns the image as a jpeg-encoded []byte
func (c *Canva) ToBytes() ([]byte, error) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, c.Img(), nil); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
