package clustering

import "testing"

func TestVectorEqual(t *testing.T) {
	tests := []struct {
		name     string
		a        Vector
		b        Vector
		expected bo