package clustering

import (
	"errors"
	"testing"
)

func TestCentroid(t *testing.T) {
	tests := []struct {
		name          string
		vectorLength  int
		data          []Vector
		expected      Vector
		expectedError error
	}{
		{
			name:          "Emp