package training

// Trainee defines the behaviour something needs to provide if it can be trained.
// Examples include rbf.Network and rbf.Node.
type Trainee interface {
	Calculate(inp