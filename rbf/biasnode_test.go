package rbf

import (
	"reflect"
	"testing"
)

func TestBiasNode(t *testing.T) {
	tests := []struct {
		name     string
		c