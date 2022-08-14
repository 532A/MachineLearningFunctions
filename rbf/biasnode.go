package rbf

// NewBias creates a bias node with the specified number of outputs.
func NewBias(count int) (b Bias) {
	b.Outputs = make([]float64, count)
	for i := 0; i < len(b.Outputs); i++ {
		b.Output