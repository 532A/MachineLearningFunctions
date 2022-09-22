package rbf

import (
	"reflect"
	"testing"
)

func TestBiasNode(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		input    []float64
		expected []float64
	}{
		{
			name:     "get the same output regardless of input (1)",
			count:    1,
			input:    []float64{6.0},
			expected: []float64{1.0},
		},
		{
			name:     "get the same output regardless of input (2)",
			count:    1,
			input:    []float64{12.0},
			expected: []float64{1.0},
		},
		{
			name:     "changing the count increases the number of output values",
			count:    3,
			input:   