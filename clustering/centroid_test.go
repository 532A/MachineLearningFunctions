package clustering

import (
	"errors"
	"testing"
)

func TestCentroid(t *testing.T) {
	tests := []struct {
		name          string
		vectorLength  int
		data          []Vector
		expected      Vector
		expectedError error
	}{
		{
			name:          "Empty",
			vectorLength:  3,
			data:          []Vector{},
			expected:      Vector{0, 0, 0},
			expectedError: errors.New("centroid: no data provided"),
		},
		{
			name:         "All ones",
			vectorLength: 3,
			data: []Vector{
				Vector{1, 1, 1},
				Vector{1, 1, 1},
				Vector{1, 1, 1},
				Vector{1, 1, 1},
			},
			expected: Vector{1, 1, 1},
		},
		{
			name:         "Average of 1 and 2 is 1.5",
			vectorLength: 3,
			data: []Vector{
				Vector{1, 2, 1},
				Vector{2, 1, 2},
			},
			expected: Vect