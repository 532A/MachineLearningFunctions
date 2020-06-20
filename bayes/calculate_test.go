package bayes

import "testing"

func TestProbability(t *testing.T) {
	categoryDog, categoryCat := Category(1), Category(2)

	categoryBrownHair, categoryBlondeHair := Category(1), Category(2)
	categoryBlueEyes, categoryBrownEyes