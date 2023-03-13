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

// Node defines a node in an RBF network which uses a distance function to calculate distance.
type Node struct {
	InputWeights []float64
	Centroid     []float64
	// RBF function parameters.
	Width         float64
	OutputWeights []float64
}

func (n Node) String() string {
	b, err := json.Marshal(n)
	if err != nil {
		return fmt.Sprintf("rbf.Node: error marshalling to JSON: %v", err)
	}
	return string(b)
}

// Calculate the output of the node.
func (n *Node) Calculate(input []float64) (op []float64, err error) {
	if len(n.InputWeights) != len(input) {
		err = fmt.Errorf("rbf: the input vector has a length of %d values and should have the same number of input weights, but we have %d node input we