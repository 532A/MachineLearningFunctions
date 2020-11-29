package calculus

// TangentSlope returns the slope of the line at point x by measuring the slope from x - d to x.
func TangentSlope(x, d float64, f func(x float64) (y float64)) Li