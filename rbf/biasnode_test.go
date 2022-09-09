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
	}{
		{
			name:     "get the same output regardless of input (1)",
			count:    1,
			input:    []