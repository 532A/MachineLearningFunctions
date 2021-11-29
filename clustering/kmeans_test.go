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
		/