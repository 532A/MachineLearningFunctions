package bayes

import "testing"

func TestProbability(t *testing.T) {
	categoryDog, categoryCat := Category(1), Category(2)

	categoryBrownHair, categoryBlondeHair := Category(1), Category(2)
	categoryBlueEyes, categoryBrownEyes := Category(3), Category(4)

	tests := []struct {
		n