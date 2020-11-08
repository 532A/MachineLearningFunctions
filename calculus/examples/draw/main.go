
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