package rbf

import "testing"

func TestGaussian(t *testing.T) {
	actual := NewGaussian(3.0, 2.0, 6.0)(2.0)
	if actual != 3.0 {
		t.Errorf("expected the peak of the curve to be at 2.0 and the value to be 3.0, but got %v", actual)
	}
}

func TestGaussianVectorSuccess(t *testing.T) {
	center := []float64{3.0, 5.0}
	f := NewGaussianVector(6.0, center, 1.0)
	actual, err := f([]float64{3.0, 5.0