package distance

import "testing"

func TestErrors(t *testing.T) {
	tests := []struct {
		name     string
		p        []float64
		q        []float64
		expected error
	}{
		{
			name:     "Zero length vectors",
			p:        []float64{},
			q:        []float64{},
			expected: ErrZeroLengthVector,
		},
		{
			name:     "Mismatched lengths",
			p:        []float64{-1},
			q:        []float64{3, 3},
			expected: 