package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"gonum.org/v1/plot/plotter"

	"github.com/a-h/ml/rbf"
	"gonum.org/v1/plot"
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

	p.Title.Text = "Gaussian and Ricker Wavelet RBFs"
	p.X.Min = 0
	p.X.Padding = 0
	p.X.Label.Text = "X"
	p.Y.Min = 0
	p.Y.Padding = 0
	p.Y.Label.Text =