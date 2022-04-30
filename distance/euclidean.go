package distance

import (
	"math"
)

// Euclidean distance between two vectors.
func Euclidean(p []float64, q []float64) (d float64, err error) {
	if err = valid