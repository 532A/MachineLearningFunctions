
package main

import (
	"fmt"
	"image"
	"image/color/palette"
	imgdraw "image/draw"
	"image/gif"
	"os"

	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot/vg/draw"

	"github.com/a-h/ml/calculus"
	"gonum.org/v1/plot/vg/vgimg"

	"gonum.org/v1/plot/plotter"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
)

func main() {
	file, err := os.Create("op.gif")
	if err != nil {
		fmt.Println("Failed to create op.gif file:", err)
		return
	}

	// Draw the y=x*x graph points.
	f := func(x float64) (y float64) { return x * x }

	var fys plotter.XYs
	step := 0.1
	minXValue, maxXValue := -6.0, 6.0
	for i := minXValue; i < maxXValue; i += step {
		fys = append(fys, xy{
			X: float64(i),
			Y: f(float64(i)),
		})
	}

	var minYValue, maxYValue float64
	for i, v := range fys {
		if v.Y < minYValue || i == 0 {
			minYValue = v.Y
		}
		if v.Y > maxYValue || i == 0 {
			maxYValue = v.Y
		}
	}

	// Now draw the tangent line and add it to the gif.
	width := vg.Centimeter * 40
	height := vg.Centimeter * 30
	dpi := 96.0
	imgSize := image.Rect(0, 0, int(width.Dots(dpi)), int(height.Dots(dpi)))
	frames := 12
	anim := gif.GIF{LoopCount: -1}
	delay := 10 // 100ms

	step = (maxXValue - minXValue) / float64(frames)
	for i := minXValue; i < maxXValue; i += step {
		// Calculate the tangent slope.
		l := calculus.TangentSlope(i, 0.5, f)

		// Always draw the function output.
		var line plotter.XYs
		for x := minXValue; x < maxXValue; x++ {
			y := l.Y(x)
			if y >= 0 {
				line = append(line, xy{
					X: x,