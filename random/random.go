package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Float64 returns a random float64 between min and max.
func Float64(min, max float64) float64 {
	return (max-min)*rand.Float64() + min
}

// Float64Vector returns an array of float64 values min and max.
func Float64Vector(min, max float64, size int) []float64 {
	op :=