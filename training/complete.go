
package training

import (
	"context"
	"fmt"

	"github.com/a-h/ml/distance"
)

// Complete executes a training session against the trainee t, using data d, algorithm a and
// the distance function dist to calculate the distance between the expected and received values.
// Stoppers can be provided to limit the training to a time period, maximum number of iterations,
// error etc.
func Complete(t Trainee, d []Data, a Algorithm, dist distance.Function, stoppers ...Stopper) (iterations int, err error) {
	for {
		e := func() (e float64, err error) {
			return evaluateTrainee(t, d, dist)
		}
		updatedMemory, err := a.Next(e)
		if err != nil {
			return iterations, fmt.Errorf("training.Complete: error at iteration %v: %v", iterations, err)
		}
		t.SetMemory(updatedMemory)

		iterations++
		if shouldStop(iterations, a.BestError(), stoppers) {
			break
		}
	}