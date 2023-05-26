
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"os/signal"
	"time"

	"github.com/a-h/ml/distance"
	"github.com/a-h/ml/projection"

	"github.com/a-h/ml/training"

	"github.com/a-h/ml/rbf"
)

func main() {
	ideal := []training.Data{
		{
			Input:    []float64{0, 0},
			Expected: []float64{0},
		},
		{
			Input:    []float64{0, 1},
			Expected: []float64{1},
		},
		{
			Input:    []float64{1, 0},
			Expected: []float64{1},
		},
		{
			Input:    []float64{1, 1},
			Expected: []float64{0},
		},
	}

	network, err := rbf.NewNetwork(
		rbf.NewNode(2, 1),