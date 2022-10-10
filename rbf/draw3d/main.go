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

	anim