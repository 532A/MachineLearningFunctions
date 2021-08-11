
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/a-h/ml/clustering"
	"github.com/a-h/ml/distance"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	p, err := plot.New()
	if err != nil {
		fmt.Println("Error creating Plot: ", err)
		os.Exit(-1)
	}

	p.Title.Text = "KMeans"
	p.X.Min = 0
	p.X.Padding = 0
	p.X.Label.Text = "X"
	p.Y.Min = 0
	p.Y.Padding = 0
	p.Y.Label.Text = "Y"

	// Create some random data and assign to n clusters.
	data := random2DVectors(50)
	n := 3
	assignment, err := clustering.KMeans(data, n, distance.Euclidean)
	if err != nil {
		fmt.Println("Error clustering data: ", err)
		os.Exit(-1)
	}

	// Get the clusters.
	clusters, err := clustering.Assign(data, assignment)
	if err != nil {
		fmt.Println("Error assigning data to clusters: ", err)
		os.Exit(-1)