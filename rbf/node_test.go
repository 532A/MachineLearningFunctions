package rbf

import (
	"reflect"
	"testing"
)

func TestNode(t *testing.T) {
	tests := []struct {
		name     string
		node     *Node
		input    []float64
		e