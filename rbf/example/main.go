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
		fmt.Println("Error creating Plot: ",