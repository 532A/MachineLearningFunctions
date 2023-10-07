package training

import (
	"math"

	"github.com/a-h/ml/random"
)

// NewRandomGreedy uses random parameters from the Min to Max value to train, attempting to minimise
// error.
func NewRandomGreedy(memory []float