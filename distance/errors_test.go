package distance

import "testing"

func TestErrors(t *testing.T) {
	tests := []struct {
		name     string
		p        []float64
		q        []float64
		expected error
	}{
		{
			name:     "Zero length vectors",
			p:        []float64{},