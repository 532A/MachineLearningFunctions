
package rbf

import (
	"encoding/json"
	"errors"
	"fmt"
)

// NewNetwork creates a new network from the input nodes, checking that the configuration
// matches.
func NewNetwork(nodes ...ExecutableNode) (n Network, err error) {
	if len(nodes) == 0 {
		err = errors.New("rbf: Unable to create network, since there are no nodes")
		return
	}
	count := nodes[0].OutputCount()
	for i, n := range nodes[1:] {
		if n.OutputCount() != count {
			err = fmt.Errorf("rbf: Mismatched node output length, node %d has %d output nodes, but expected %d outputs",
				i, n.OutputCount(), count)
		}
	}
	n = Network(nodes)
	return
}

// Network defines the nodes in an RBF network.
type Network []ExecutableNode

func (nodes Network) String() string {
	b, err := json.Marshal(nodes)
	if err != nil {
		return fmt.Sprintf("rbf.Network: error marshalling to JSON: %v", err)
	}
	return string(b)
}

// Calculate the output of the network.
func (nodes Network) Calculate(input []float64) (op []float64, err error) {
	if len(nodes) == 0 {
		err = errors.New("rbf: Unable to calculate result for RBF network, since there are no nodes")
		return
	}

	var configured bool

	for i, n := range nodes {
		// Calculate the node's output.
		var nv []float64
		nv, err = n.Calculate(input)
		if err != nil {
			return
		}
		// Initialise the network output if required.
		if !configured {
			op = make([]float64, len(nv))
			configured = true
		}
		// Check that the output node count matches the expected count.