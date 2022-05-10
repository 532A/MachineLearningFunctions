package distance

import "testing"

func TestEuclidean(t *testing.T) {
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
			expected: 5,
		},
		{
			name:     "Pythagoran triangle shifted left",
			p:        []float64{-1, 0},
			q:        []float64{3, 3},
			expected: 5,
		},
		{
		