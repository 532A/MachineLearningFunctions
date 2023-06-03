
package training

import (
	"context"
	"testing"
)

func TestCompleteUsesAllTrainingData(tt *testing.T) {
	t := &traineeMock{}
	d := []Data{
		{
			Input:    []float64{0},
			Expected: []float64{0},
		},
		{
			Input:    []float64{1},
			Expected: []float64{1},
		},
	}
	a := &algorithmMock{}

	distanceCalled := 0
	dist := func(p []float64, q []float64) (d float64, err error) {
		distanceCalled++
		return 0.0, nil
	}

	iterationCount := 2

	Complete(t, d, a, dist, StopAfterXIterations(iterationCount))
