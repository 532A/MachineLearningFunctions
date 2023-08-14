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
	// Tr