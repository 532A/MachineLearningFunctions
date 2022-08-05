
package projection

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/a-h/raster"
	"golang.org/x/image/colornames"
)

// A Projection can write out a graph.
type Projection struct {
	// The size of the image to draw
	Width, Height int
	// The minimum x and y values.
	MinimumValue float64
	// The maximum x and y values.
	MaximumValue float64
	// The number of cells to render.
	Cells int
	// The function used to produce the z value.
	Function func(x, y float64) float64
	// The angle of the projection
	Angle    float64
	cosAngle float64
	sinAngle float64
	// The Z scale of the projection
	Scale           float64
	BackgroundColor color.RGBA
	LineColor       color.RGBA
	FillColor       color.RGBA
}

func New(min, max float64, function func(x, y float64) float64, width, height int, angle float64) Projection {
	rads := math.Pi / (float64(180) / angle)
	return Projection{
		Width:           width,
		Height:          height,
		MinimumValue:    min,
		MaximumValue:    max,
		Function:        function,
		Cells:           100,
		Angle:           rads,
		cosAngle:        math.Cos(rads),
		sinAngle:        math.Sin(rads),
		Scale:           float64(height) * 0.4,
		BackgroundColor: color.RGBA{},
		LineColor:       colornames.Blue,
		FillColor:       colornames.Lightblue,
	}
}

func (p Projection) Draw(img draw.Image) {
	r := p.MaximumValue - p.MinimumValue
	for i := 0; i < p.Cells; i++ {
		for j := 0; j < p.Cells; j++ {
			ax, ay, ok := corner(r, p.Width, p.Height, p.cosAngle, p.sinAngle, p.Scale, i+1, j, p.Cells, p.Function)
			if !ok {
				continue
			}
			bx, by, ok := corner(r, p.Width, p.Height, p.cosAngle, p.sinAngle, p.Scale, i, j, p.Cells, p.Function)
			if !ok {
				continue
			}
			cx, cy, ok := corner(r, p.Width, p.Height, p.cosAngle, p.sinAngle, p.Scale, i, j+1, p.Cells, p.Function)
			if !ok {
				continue
			}
			dx, dy, ok := corner(r, p.Width, p.Height, p.cosAngle, p.sinAngle, p.Scale, i+1, j+1, p.Cells, p.Function)
			if !ok {
				continue