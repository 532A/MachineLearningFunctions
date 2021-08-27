package clustering

import (
	"errors"
	"math/rand"
	"time"

	"github.com/a-h/ml/distance"
)

// KMeans cluster the input vectors into n clusters u