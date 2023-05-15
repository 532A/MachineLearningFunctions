package rbf

import (
	"reflect"
	"testing"
)

func TestNode(t *testing.T) {
	tests := []struct {
		name     string
		node     *Node
		input    []float64
		expected []float64
	}{
		{
			name: "returns RBF distance from the input",
			node: &Node{
				Width:         0.5,
				Centroid:      []float64{1.0},
				InputWeights:  []float64{1.0},
				OutputWeights: []float64{1.0},
			},
			input:    []float64{1.0},
			expected: []float64{1.0},
		},
	}

	for _, test := range tests {
		actual, err := test.node.Calculate(test.input)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", test.name, err)
			continue
		}
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%s: for input %v, expected output %v, but got %v", test.name,
				test.input, test.expected, actual)
		}
		if test.node.OutputCount() != len(test.node.OutputWeights) {
			t.Errorf("%s: expected output count of %d, got %d", test.name, len(test.node.OutputWeights), test.node.OutputCo