package rbf

// Function is a radial basis function, for use in an RBF network.
// Examples are rbf.NewGaussian and rbf.RickerWavelet.
type Function func(v flo