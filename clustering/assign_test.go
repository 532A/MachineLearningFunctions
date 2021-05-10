
package clustering

import (
	"testing"
)

func TestAssign(t *testing.T) {
	tests := []struct {
		name       string
		input      []Vector
		assignment []int
		expected   []Cluster
	}{
		{
			name: "All to one",
			input: []Vector{
				Vector{1, 2, 3},
				Vector{4, 5, 6},
			},
			assignment: []int{1, 1},
			expected: []Cluster{