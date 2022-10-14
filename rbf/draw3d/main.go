package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"

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

	anim := gif.GIF{LoopCount: 128}

	delay := 1 // 10ms

	deviationFrom, deviationTo := 0.0, 10.0
	deviationStep := (deviationFrom - deviationTo) / float64(anim.LoopCount)

	for i := 0; i < anim.LoopCount; i++ {
		img := image.NewPaletted(imgSize, palett