package rbf

import (
	"encoding/json"
	"fmt"

	"github.com/a-h/ml/random"
)

// NewNode returns a new Node.
func NewNode(inputCount int, outputCount int) *Node {
	return &Node{
		InputWeights:  ra