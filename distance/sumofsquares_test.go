package distance

import "testing"

func TestSumOfSquares(t *testing.T) {
	tests := []struct {
		name     string
		p        []float64
		q        []float64
		expected float64
	}{
		{
			name:     "Equal",
			p:        []float64{1, 2, 3, 4, 5},
			q:        []float64{1, 2, 3, 