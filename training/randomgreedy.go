package training

import (
	"math"

	"github.com/a-h/ml/random"
)

// NewRandomGreedy uses random parameters from the Min to Max value to train, attempting to minimise
// error.
func NewRandomGreedy(memory []float64) *RandomGreedy {
	min, max := -10.0, 10.0
	return &RandomGreedy{
		Min:     min,
		Max:     max,
		current: memory,
		best:    nil,
		e:       math.MaxFloat64,