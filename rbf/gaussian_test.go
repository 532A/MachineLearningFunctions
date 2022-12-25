package rbf

import "testing"

func TestGaussian(t *testing.T) {
	actual := NewGaussian(3.0, 2.0, 6.0)(2.0)
	if actual != 3.0 {
		t.Errorf("expected