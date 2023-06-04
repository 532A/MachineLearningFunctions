
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"

	"github.com/a-h/ml/training"

	"github.com/a-h/ml/projection"
	"github.com/a-h/ml/rbf"
)

var palette = []color.Color{color.White, color.Black}

func main() {
	file, err := os.Create("op.gif")
	if err != nil {
		fmt.Println("Failed to create op.gif file:", err)
		return
	}

	imgSize := image.Rect(0, 0, 500, 500)

	anim := gif.GIF{LoopCount: 32}

	delay := 1 // 10ms

	// Make some "terrain" out of 3D gaussian functions.
	hills := []rbf.VectorFunction{
		rbf.NewGaussianVector(1.0, []float64{-0.5, -0.25}, 0.25),
		rbf.NewGaussianVector(0.75, []float64{0.25, -0.25}, 0.5),
		rbf.NewGaussianVector(0.5, []float64{-0.25, 0.25}, 0.4),
		rbf.NewGaussianVector(0.8, []float64{0.5, 0.25}, 0.3),
	}

	f := func(x, y float64) (z float64) {
		for _, f := range hills {
			r, err := f([]float64{x, y})
			if err != nil {
				panic(err)
			}