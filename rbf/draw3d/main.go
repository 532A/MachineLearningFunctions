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
	