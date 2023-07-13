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
		current:  memory,
		e:        math.MaxFloat64,
		velocity: velocity,
	}
	hc.movements = []float64{
		-acceleration,     // Go back
		-1 / acceleration, // Go back slower
		0,                 // Stop
		1 / acceleration,  // Go forward slower
		acceleration,      // Go forward
	}
	return hc
}

// HillClimbing is a training algorithm which uses random parameters from the Min to Max value to train,
// attempting to minimise error.
type HillClimbing struct {
	current []float64
	// best memory and error recorded during training.
	best []float64
	e    float64
	// Min and Max values for memory values.
	Min, Max float64
	// The amount to move in each direction.
	velocity float64
	// movements (forward, back, stay still)
	movements []float64
}

// Next records the error from the previous set of values, and returns a 