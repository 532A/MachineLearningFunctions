package rbf

import (
	"reflect"
	"testing"
)

func TestBiasNode(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		input    []float64
		expected []float64
