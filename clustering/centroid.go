
package clustering

import "errors"

// Centroid calculates the average position of the members in the cluster.
func Centroid(data []Vector) (c Vector, err error) {
	if data == nil || len(data) == 0 {
		err = errors.New("centroid: no data provided")
		return
	}

	// Initialize array.
	c = make([]float64, len(data[0]))

	// Sum the data.