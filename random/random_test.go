
package random

import "testing"

func TestFloat64(t *testing.T) {
	tests := []struct {
		min, max float64
	}{
		{
			min: -1.0,
			max: 1.0,
		},
		{
			min: 5.0,
			max: 10.0,
		},