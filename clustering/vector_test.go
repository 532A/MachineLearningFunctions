package clustering

import "testing"

func TestVectorEqual(t *testing.T) {
	tests := []struct {
		name     string
		a        Vector
		b        Vector
		expected bool
	}{
		{
			name:     "All ones",
			a:        Vector{1, 1, 1},
			b:        Vector{1, 1, 1},
			expected: true,
		},
	