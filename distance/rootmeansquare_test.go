package distance

import (
	"math"
	"testing"
)

func TestRootMeanSquare(t *testing.T) {
	tests := []struct {
		name     string
		p        []float64
		q        []float64
		expected float64
	}{
		{
			name:     "Equal",
			p:        []float64{1, 2, 3, 4, 5},
			q:        []float64{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "One above",
			p:        []float64{1, 2, 3, 4, 5},
			q:        []float64{2, 3, 4, 5, 6},
			expected: 1, // Square root of average squared error is 1.
		},
		{
			name:     "Two below",
			p:        []float64{2, 3, 4, 5, 6},
			q:        []float64{0, 1, 2, 3, 4},
			expected: 2, // Square root