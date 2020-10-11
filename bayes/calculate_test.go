package bayes

import "testing"

func TestProbability(t *testing.T) {
	categoryDog, categoryCat := Category(1), Category(2)

	categoryBrownHair, categoryBlondeHair := Category(1), Category(2)
	categoryBlueEyes, categoryBrownEyes := Category(3), Category(4)

	tests := []struct {
		name     string
		data     Data
		is       Category
		given    []Category
		expected float64
	}{
		{
			name: "dog or cat",
			data: []Datum{
				NewDatum("Felix", categoryCat),
				NewDatum("Shep", categoryDog),
			},
			is:       categoryCat,
			given:    []Category{},
			expected: 0.5,
		},
		{
			name: "brown or blue eyes",
			data: []Datum{
				NewDatum("John", categoryBlueEyes, categoryBlondeHair),
		