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
			name:  