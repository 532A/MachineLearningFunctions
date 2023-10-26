
package training

import (
	"testing"

	"github.com/a-h/ml/array"
)

func TestRandomGreedy(t *testing.T) {
	a := []float64{1.0, 2.0, 3.0, 4.0}
	rg := NewRandomGreedy(a)
	rg.Min = -100
	rg.Max = 100

	ev := func() (e float64, err error) {
		return 123.0, nil
	}
	b, err := rg.Next(ev)
	if err != nil {
		t.Fatalf("unexpected error calling Next: %v", err)
	}

	if array.EqFloat64(a, b) {