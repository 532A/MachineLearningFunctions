package bayes

// Category of the data.
type Category int

// NewDatum creates new data.
func NewDatum(v interface{}, categories ...Category) Datum {
	d := Datum{
		Value:      v,
		Categories: make(map[Category]interface{}, len(categories)),
	}
	for _, c := range categories {
		d.Categories[c] = struct{}{}
	}
	return d
}

// Datum and its associated categories.
type Datum struct {
	Value      interface{}
	Categories map[Category]interface{}
}

// Data is a collection.
type Data []Datum

// Probability of the data being in a 