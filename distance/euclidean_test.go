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
			p: 