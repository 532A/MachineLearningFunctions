package rbf

import "testing"

func TestGaussian(t *testing.T) {
	actual := NewGaussian(3.0, 2.0, 6.0)(2.0)
	if actual != 3.0 {
		t.Errorf("expected the peak of the curve to be at 2.0 and the value to be 3.0, but got %v", actual)
	}
}

func TestGaussianVectorSuccess(t *testing.T) {
	center := []flo