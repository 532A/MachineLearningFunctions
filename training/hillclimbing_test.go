package training

import (
	"errors"
	"math"
	"reflect"
	"testing"
)

func TestHillClimbing(t *testing.T) {
	// Start with 1*1
	memory := []float64{1, 1}
	desiredOutput := 25.0
	// Train finding two numbers which can be multiplied together to equal 25.
	toTrain := func() (e float64, err error) {
		result := memory[0] * memory[1]
		e = math.Pow(result-desiredOutput, 2)
		return
	}
	// Train the algorithm, but step by 1 each time, so it's easy to test and know that we'll
	// get the desired factors of 5*5 in just a few steps.
	hc := NewHillClimbing(memory, 1, 1)
	for i := 1; i < 5; i++ {
		updatedMemory, err := hc.Next(toTrain)
		if err != nil {
			t