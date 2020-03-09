package bayes

// Category of the data.
type Category int

// NewDatum creates new data.
func NewDatum(v interface{}, categories ...Category) Datum {
	d := Datum{
		Va