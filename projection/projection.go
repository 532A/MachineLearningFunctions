
package projection

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/a-h/raster"
	"golang.org/x/image/colornames"
)

// A Projection can write out a graph.
type Projection struct {
	// The size of the image to draw
	Width, Height int
	// The minimum x and y values.
	MinimumValue float64
	// The maximum x and y values.
	MaximumValue float64