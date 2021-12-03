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
		