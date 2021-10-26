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

	// Assign data to random clusters, but make sure every cluster has something in it.
	assignment = make([]int, len(data))

	r := rand.New(rand.NewSource(time.Now().Unix()))
	assigned := map[int]interface{}{}
	for i := 0; i < n; i++ {
		for {
			to := r.Intn(len(data))
			if _, ok := assigned[to]; !ok {
				assignment[to] = i
				assigned[to] = true
				break
			}
		}
	}
	for i := 0; i < len(data); i++ {
		if _, ok := assigned[i]; !ok {
			assignment[i] = r.Intn(n)
		}
	}

	// Create the centroids array once.
	centroids := make([]Vector, n)

	var done bool
	for {
		// Exit when done processing.
		if done {
			break
		}
		// Calculate / recalculate centroids.
		err = Centroids(data, n, assignment, &centroids)
		if err != nil {
			return assignment, err
		}
		done = true
		for i, v := range data {
			currentAssignmentIndex := assignment[i]
			newAss