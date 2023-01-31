package rbf

import (
	"reflect"
	"testing"
)

func TestNetworkMemory(t *testing.T) {
	a, err := NewNetwork(
		