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
			InputWeights:  []float64{3.0, 4.0},
			OutputWeights: []float64{5.0, 6.0},
		},
		&Node{
			Width:         1.0,
			Centroid:      []float64{1.0, 2.0}, // Centroid doesn't change.
			InputWeights:  []float64{3.0, 4.0},
			OutputWeights: []float64{3.0, 3.0},
		},
	)
	if err != nil {
		t.Fatalf("Failed to create network a: %v", err)
	}
	b, err := NewNetwork(
		&Node{
			Width:         -1,
			Centroid:      []float64{1.0, 2.0},
			InputWeights:  []float64{-1, -1},
			OutputWeights: []float64{-1, -1},
		},
		&Node{
			Width:         -1,
			Centroid:      []float64{1.0, 2.0}, // Centroid doesn't change.
			InputWeights:  []float64{-1, -1},
			OutputWeights: []float64{-1, -1},
		},
	)
	if err != nil {
		t.Fatalf("Failed to create network b: %v", err)
	}
	b.SetMemo