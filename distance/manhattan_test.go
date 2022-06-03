package distance

import "testing"

func TestManhattan(t *testing.T) {
	tests := []struct {
		name     string
		p        []float64
		q        []float64
		expected float64
	}{
		{
			name:     "Pythagoran triangle from origin",
			p:        []float64{0, 0},
			q:        []float64{4, 3},
			expected: (4 - 0) + (3 - 0),
		},
		{
			name:     "Pythagoran triangle 2",
			p:        []float64{4, 3},
			q:        []float64{0, 0},
			expected: 4 + 3,
		},
		{
			name:     "Equal values",
			p:        []float64{-1, 6},
			q:        []float64{-1, 6},
			expected: 0 + 0,
		},
		{
			name:     "Negative values",