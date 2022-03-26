package distance

import (
	"errors"
)

// ErrMismatchedVectorLengths is an error for when an operation cannot be carried out due to different sizes
// of vectors being used, e.g. 2 vectors, one with length 2 and t