package clustering

import (
	"math/rand"
	"testing"

	"github.com/a-h/ml/distance"
)

func TestKMeans(t *testing.T) {
	tests := []struct {
		name  string
		input []Vector
		n     int
		// Each array holds the expected contents of a cluster.
		// For example, expected {0, 1} means that input values
		// input[0] and input[1] should be present together in a
		// cluster.
		expected [][]int
	}{
		{
			name: "Single output",
			input: []Vector{
				{0, 0, 0},
				{1, 1, 1},
			},
			n: 1,
			expected: [][]int{
				{0, 1},
			},
		},
		{
			name: "Two inputs, two outputs",
			input: []Vector{
				{0, 0, 0},
				{1, 1, 1},
			},
			n: 2,
			expected: [][]int{
				{0},
				{1},
			},
		},
		{
			name: "Left vs Right",
			input: []Vector{
				{-12, 0},
				{-16, 0},
				{12, 0},
				{16, 0},
			},
			n: 2,
			expected: [][]int{
				{0, 1},
				{2, 3},
			},
		},
		{
			name: "Up vs Down",
			input: []Vector{
				{3, -100},
				{40, -200},
				{16, 100},
				{10, 200},
			},
			n: 2,
			expected: [][]int{
				{0, 1},
				{2, 3},
			},
		},
	}

	for _, test := range tests {
		assignment, err := KMeans(test.input, test.n, distance.Euclidean)
		if err != nil {
			t.Fatalf("%s: unexpected error: %v", test.name, err)
		}
		if len(assignment) != len(test.input) {
			t.Errorf("%s: expected assignment length %d to equal the amount of data %d", test.name, len(assignment), len(test.input))
		}

		// Clusters are assigned randomly and may be in any order.
		allActualClusters := getClusters(test.input, assignment)

		for _, expectedCluster := range test.expected {
			// Find the first cluster that contains the expected item. The rest of the items should also be
			// there.
			firstExpected := test.input[expectedCluster[0]]

			var allOthersShouldBeFoundIn int
			for i, a := range allActualClusters {
			out:
				for _, vv := range a {
					if firstE