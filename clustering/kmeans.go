package clustering

import (
	"errors"
	"math/rand"
	"time"

	"github.com/a-h/ml/distance"
)

// KMeans cluster the input vectors into n clusters using the distance function d.
func KMeans(data []Vector, n int, d distance.Function) (assignment []int, err error) {
	if n <= 0 {
		return nil, errors.New("KMeans: n cannot be less than or equal to zero")
	}
	if data == nil {
		return nil, errors.New("KMeans: data cannot be nil")
	}
	if len(data) == 0 {
		return nil, errors.New("KMeans: data cannot be empty")
	}

	//