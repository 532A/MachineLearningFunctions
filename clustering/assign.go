package clustering

import (
	"errors"
)

// Assign vectors to their clusters.
func Assign(data []Vector, assignment []int) ([]Cluster, error) {
	if len(assignment) != len(data) {
		return nil, errors.New("assignment must equal the amount of input")
	}
	clusters := 0
	for _, v := range assignment {
		if v > clusters {
			clusters = v
		}
	}
	op := make([]Cluster, clusters+1)
	for i, a 