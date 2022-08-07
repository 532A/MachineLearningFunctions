package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Float64 returns a random float64 between min and max.
func Float64(min, 