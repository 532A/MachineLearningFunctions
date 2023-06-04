
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

	expectedCalculations := len(d) * iterationCount
	if t.calculateCalled != expectedCalculations {
		tt.Errorf("expected %d training calculations, but %d were carried out", expectedCalculations, t.calculateCalled)
	}
	expectedDistanceCalculations := expectedCalculations
	if distanceCalled != expectedDistanceCalculations {
		tt.Errorf("expected %d distance calculations, but %d were carried out", expectedDistanceCalculations, distanceCalled)
	}
	expectedNextCalled := len(d)
	if a.nextCalled != expectedNextCalled {
		tt.Errorf("expected %d next called, but %d were carried out", expectedNextCalled, a.nextCalled)
	}
	expectedSetMemoryCalled := len(d)
	if t.setMemoryCalled != expectedSetMemoryCalled {
		tt.Errorf("expected %d set memory called, but %d were carried out", expectedSetMemoryCalled, t.setMemoryCalled)
	}
	if a.bestErrorCalled != 2 {
		tt.Errorf("expected bestMemory to be called twice, but was called %d", a.bestErrorCalled)
	}
}

type traineeMock struct {
	calculateCalled     int
	getMemorySizeCalled int
	getMemoryCalled     int
	setMemoryCalled     int
	memory              []float64
}

func (tm *traineeMock) Calculate(input []float64) (output []float64, err error) {
	tm.calculateCalled++
	return input, nil
}
func (tm *traineeMock) GetMemorySize() int {
	tm.getMemorySizeCalled++
	return len(tm.memory)
}
func (tm *traineeMock) GetMemory() []float64 {
	tm.getMemoryCalled++
	return tm.memory
}
func (tm *traineeMock) SetMemory(m []float64) {
	tm.setMemoryCalled++
	tm.memory = m
}

type algorithmMock struct {
	nextCalled       int
	bestErrorCalled  int
	bestMemoryCalled int
	memory           []float64
	e                float64
}

func (am *algorithmMock) Next(ev Evaluator) (updatedMemory []float64, err error) {
	am.nextCalled++
	am.e, err = ev()