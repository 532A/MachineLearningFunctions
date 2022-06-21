package distance

import (
	"math"
	"testing"
)

func TestRootMeanSquare(t *testing.T) {
	tests := []struct {
		name     string
		p        []float64
		q        []float64
		expected float64
	}{
		{
			name:     "E