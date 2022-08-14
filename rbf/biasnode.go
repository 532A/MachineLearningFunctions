package rbf

// NewBias creates a bias node with the specified number of outputs.
func NewBias(count int) (b Bias) {
	b.Outputs = make([]float64, co