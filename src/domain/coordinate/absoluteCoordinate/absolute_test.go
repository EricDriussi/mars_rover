package absoluteCoordinate_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualsBasedOnValues(t *testing.T) {
	aCoordinate := absoluteCoordinate.From(1, 1)
	anEqualCoordinate := absoluteCoordinate.From(1, 1)

	areTheSame := aCoordinate.Equals(anEqualCoordinate)
	assert.True(t, areTheSame)
}

func TestNotEqualsBasedOnValues(t *testing.T) {
	aCoordinate := absoluteCoordinate.From(1, 1)
	testCases := []struct {
		name                string
		differentCoordinate *AbsoluteCoordinate
	}{
		{
			name:                "neither X nor Y can differ",
			differentCoordinate: absoluteCoordinate.From(2, 2),
		},
		{
			name:                "X cannot differ",
			differentCoordinate: absoluteCoordinate.From(2, 1),
		},
		{
			name:                "Y cannot differ",
			differentCoordinate: absoluteCoordinate.From(1, 2),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			areTheSame := aCoordinate.Equals(testCase.differentCoordinate)
			assert.False(t, areTheSame)
		})
	}
}
