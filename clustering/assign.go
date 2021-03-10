package clustering

import (
	"errors"
)

// Assign vectors to their clusters.
func Assign(data []Vector, assignment []int) ([]Cluste