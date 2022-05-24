
package distance

import (
	"math"
)

// Manhattan distance between two vectors.
func Manhattan(p []float64, q []float64) (d float64, err error) {
	if err = validateInputs(p, q); err != nil {
		return