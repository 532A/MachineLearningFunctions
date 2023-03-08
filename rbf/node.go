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
	return str