package rbf

import (
	"encoding/json"
	"fmt"

	"github.com/a-h/ml/random"
)

// NewNode returns a new Node.
func NewNode(inputCount int, outputCount int) *Node {
	return &Node{
		InputWeights:  random.Float64Vector(-10, 10, inputCount),
		Centroid:      random.Float64Vector(-10, 10, inputCount),
		Width:         random.Float64(-10, 10),
		OutputWeights: random.Float64Vector(-10, 10, outputCount),
	}
}

// Node defines a node in an R