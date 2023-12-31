package absoluteCoordinate_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualsBasedOnValues(t *testing.T) {
	aCoordinate := absoluteCoordinate.Build(1, 1)
	anEqualCoordinate := absoluteCoordinate.Build(1, 1)

	assert.True(t, aCoordinate.Equals(*anEqualCoordinate))
}

func TestNotEqualsBasedOnValues(t *testing.T) {
	aCoordinate := absoluteCoordinate.Build(1, 1)
	testCases := []struct {
		name                string
		differentCoordinate *AbsoluteCoordinate
	}{
		{
			name:                "both coordinates are different",
			differentCoordinate: absoluteCoordinate.Build(2, 2),
		},
		{
			name:                "X is different",
			differentCoordinate: absoluteCoordinate.Build(2, 1),
		},
		{
			name:                "Y is different",
			differentCoordinate: absoluteCoordinate.Build(1, 2),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.False(t, aCoordinate.Equals(*testCase.differentCoordinate))
		})
	}
}
