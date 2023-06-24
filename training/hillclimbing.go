package training

import (
	"math"
)

// NewHillClimbing uses random parameters from the Min to Max value to train, attempting to minimise
// error.
func NewHillClimbing(memory []float64, velocity, acceleration float64) *HillClimbing {
	min, max := -10.0, 10.0
	hc := &HillClimbing{
		Min:      min,
		Max:      max,
		current:  m