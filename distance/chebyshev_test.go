package distance

import "testing"

func TestChebyshev(t *testing.T) {
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
			expected: 4,
		},
		{
			name:     "Larger X",
			p:        []float64{32, 0},
			q:        []float64{16, 3},
			expected: 16,
		},
		{
			name:     "Larger Y",
			p:        []float64{0, -32},
			q:        []