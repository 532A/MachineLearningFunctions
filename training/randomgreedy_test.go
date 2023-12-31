
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
		t.Errorf("expected a and b to be different, but they were the same")
	}

	for i := 0; i < 10000; i++ {
		updatedMemory, err := rg.Next(func() (e float64, err error) {
			return 200.0, nil
		})
		if err != nil {
			t.Fatalf("unexpected error at iteration %d: %v", i, err)
		}
		for _, v := range updatedMemory {
			if v > 100 {
				t.Errorf("unexpected value > 100")
			}
			if v < -100 {
				t.Errorf("unexpected value < 100")
			}
		}
	}

	if rg.BestError() != 123.0 {
		t.Errorf("expected the best error of 123.0, but got %v", rg.BestError())
	}
	if !array.EqFloat64(rg.BestMemory(), a) {
		t.Errorf("expected the best memory to be %v, but got %v", a, rg.BestMemory())
	}

	rg.Next(func() (e float64, err error) {
		return 0, nil
	})
	if array.EqFloat64(rg.BestMemory(), a) {
		t.Errorf("the best memory should have been randomly generated, but got %v", rg.BestMemory())
	}
}