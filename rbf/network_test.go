package rbf

import (
	"reflect"
	"testing"
)

func TestNetworkMemory(t *testing.T) {
	a, err := NewNetwork(
		&Node{
			Width:         0.5,
			Centroid:      []