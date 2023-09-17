package training

// Trainee defines the behaviour something needs to provide if it can be trained.
// Examples include rbf.Network and rbf.Node.
type Trainee interface {
	Calculate(input []float64) (output []float64, err error)
	GetMemorySize() int
	GetMemory() []float64
	SetMemory(m []float64)
}

// A