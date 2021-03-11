package clustering

import (
	"errors"
)

// Assign vectors to their clusters.
func Assign(data []Vector, assignment []int) ([]Cluster, error) {
	if len(assignment) != len(data) {
		return nil, errors.New("assignment must equal the am