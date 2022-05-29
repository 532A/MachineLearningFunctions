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
			name:     "Pythagoran tria