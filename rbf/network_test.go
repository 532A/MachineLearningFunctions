package rbf

import (
	"reflect"
	"testing"
)

func TestNetworkMemory(t *testing.T) {
	a, err := NewNetwork(
		&Node{
			Width:         0.5,
			Centroid:      []float64{1.0, 2.0},
			InputWeights:  []float6