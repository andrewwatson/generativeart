package arts

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/andrewwatson/generativeart"
	"github.com/andrewwatson/generativeart/common"
	"github.com/fogleman/gg"
)

const (
	// minRadius        = 10
	// maxRadius        = 30
	defaultSpacing   = 30.0
	defaultAmplitude = 40.0

	minLineWidth = 3.0
	maxLineWidth = 10.0
)

type notewave struct {
	startingY  float64
	noise      *common.PerlinNoise
	spacing    float64
	wavelength float64
	amplitude  float64
	lines      bool
	noteName   string
	minRadius  float64
	maxRadius  float64
}

// wavelengths can be obtained at https://pages.mtu.edu/~suits/notefreqs.html
func NewNoteWave(startingY float64, noteName string, minRadius, maxRadius float64) *notewave {

	return &notewave{
		startingY: startingY,
		noise:     common.NewPerlinNoise(),
		spacing:   defaultSpacing,
		// wavelength: wavelength / 200,
		amplitude: defaultAmplitude,
		lines:     false,
		noteName:  noteName,
		minRadius: minRadius,
		maxRadius: maxRadius,
	}

}
func (nw *notewave) Generative(c *generativeart.Canva) string {

	wavelength, err := nw.GetWavelength(nw.noteName)
	if err != nil {
		fmt.Printf("ERR: %s\n", err.Error())
	}
	fmt.Printf("%s Wavelength: %0.6f\n", nw.noteName, wavelength)
	ctex := gg.NewContextForRGBA(c.Img())

	startingX := float64(c.Width()) * 0.05
	endingX := float64(c.Width()) * 0.95

	for i := startingX; i < endingX; i += nw.spacing {

		var lw float64
		if rand.Float64() < 0.6 {
			lw = minLineWidth
		} else {
			lw = common.RandomRangeFloat64(minLineWidth, common.RandomRangeFloat64(minLineWidth, maxLineWidth))
		}
		ctex.SetLineWidth(lw)

		radianX := gg.Radians(i*wavelength) / 80
		SinX := math.Sin(radianX)
		amplitude := SinX * nw.amplitude

		cls := c.Opts().ColorSchema()[rand.Intn(len(c.Opts().ColorSchema()))]
		ctex.SetColor(cls)

		radius := common.RandomRangeFloat64(nw.minRadius, nw.maxRadius)
		// fmt.Printf("Drawing Circle: i %0.02f radians %0.02f amp %0.02f sin %0.02f rad %0.02f\n", i, radianX, amplitude, SinX, nw.radius)
		ctex.DrawCircle(i, nw.startingY+(amplitude), radius)
		// ctex.DrawEllipse(i, nw.startingY+(amplitude))
		ctex.Stroke()
	}
	return ""
}

func (nw *notewave) GetWavelength(noteName string) (float64, error) {

	r := csv.NewReader(strings.NewReader(frequencyTable))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	wavelengthTable := make(map[string]string, 0)

	for _, r := range records {
		wavelengthTable[r[0]] = r[2]
	}

	if wl, ok := wavelengthTable[noteName]; ok {
		return strconv.ParseFloat(wl, 64)
	}
	return 0.0, nil
}

const frequencyTable = `
C0,16.35,2109.89
C#0/Db0,17.32,1991.47
D0,18.35,1879.69
D#0/Eb0,19.45,1774.2
E0,20.6,1674.62
F0,21.83,1580.63
F#0/Gb0,23.12,1491.91
G0,24.5,1408.18
G#0/Ab0,25.96,1329.14
A0,27.5,1254.55
A#0/Bb0,29.14,1184.13
B0,30.87,1117.67
C1,32.7,1054.94
C#1/Db1,34.65,995.73
D1,36.71,939.85
D#1/Eb1,38.89,887.1
E1,41.2,837.31
F1,43.65,790.31
F#1/Gb1,46.25,745.96
G1,49,704.09
G#1/Ab1,51.91,664.57
A1,55,627.27
A#1/Bb1,58.27,592.07
B1,61.74,558.84
C2,65.41,527.47
C#2/Db2,69.3,497.87
D2,73.42,469.92
D#2/Eb2,77.78,443.55
E2,82.41,418.65
F2,87.31,395.16
F#2/Gb2,92.5,372.98
G2,98,352.04
G#2/Ab2,103.83,332.29
A2,110,313.64
A#2/Bb2,116.54,296.03
B2,123.47,279.42
C3,130.81,263.74
C#3/Db3,138.59,248.93
D3,146.83,234.96
D#3/Eb3,155.56,221.77
E3,164.81,209.33
F3,174.61,197.58
F#3/Gb3,185,186.49
G3,196,176.02
G#3/Ab3,207.65,166.14
A3,220,156.82
A#3/Bb3,233.08,148.02
B3,246.94,139.71
C4,261.63,131.87
C#4/Db4,277.18,124.47
D4,293.66,117.48
D#4/Eb4,311.13,110.89
E4,329.63,104.66
F4,349.23,98.79
F#4/Gb4,369.99,93.24
G4,392,88.01
G#4/Ab4,415.3,83.07
A4,440,78.41
A#4/Bb4,466.16,74.01
B4,493.88,69.85
C5,523.25,65.93
C#5/Db5,554.37,62.23
D5,587.33,58.74
D#5/Eb5,622.25,55.44
E5,659.25,52.33
F5,698.46,49.39
F#5/Gb5,739.99,46.62
G5,783.99,44.01
G#5/Ab5,830.61,41.54
A5,880,39.2
A#5/Bb5,932.33,37
B5,987.77,34.93
C6,1046.5,32.97
C#6/Db6,1108.73,31.12
D6,1174.66,29.37
D#6/Eb6,1244.51,27.72
E6,1318.51,26.17
F6,1396.91,24.7
F#6/Gb6,1479.98,23.31
G6,1567.98,22
G#6/Ab6,1661.22,20.77
A6,1760,19.6
A#6/Bb6,1864.66,18.5
B6,1975.53,17.46
C7,2093,16.48
C#7/Db7,2217.46,15.56
D7,2349.32,14.69
D#7/Eb7,2489.02,13.86
E7,2637.02,13.08
F7,2793.83,12.35
F#7/Gb7,2959.96,11.66
G7,3135.96,11
G#7/Ab7,3322.44,10.38
A7,3520,9.8
A#7/Bb7,3729.31,9.25
B7,3951.07,8.73
C8,4186.01,8.24
C#8/Db8,4434.92,7.78
D8,4698.63,7.34
D#8/Eb8,4978.03,6.93
E8,5274.04,6.54
F8,5587.65,6.17
F#8/Gb8,5919.91,5.83
G8,6271.93,5.5
G#8/Ab8,6644.88,5.19
A8,7040,4.9
A#8/Bb8,7458.62,4.63
B8,7902.13,4.37
`
