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
	}
}

// RandomGreedy is a training algorithm which uses random parameters from the Min to Max value to train,
// attempting to minimise error.
type RandomGreedy struct {
	current []float64
	// best memory and error recorded during training.
	best []float64
	e    float64
	// Min and Max values for memory values.
	Min, Max float64
}

// Next records the error from the previous set of values, and returns a new set of values to try.
func (rg *RandomGreedy) Next(ev Evaluator) ([]float64, error) {
	e, err := ev()
	if err != nil {
		return rg.current, err
	}
	if e < rg.e {
		rg.best = rg.current
		rg.e = e
	}
	rg.current = random.Float64Vector(rg.Min, rg.Ma