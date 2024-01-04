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

func TestIsAdjacentTo(t *testing.T) {
	aCoordinate := absoluteCoordinate.Build(1, 1)
	testCases := []struct {
		name                  string
		adjacentCoordinate    *AbsoluteCoordinate
		nonAdjacentCoordinate *AbsoluteCoordinate
	}{
		{
			name:                  "X is adjacent",
			adjacentCoordinate:    absoluteCoordinate.Build(2, 1),
			nonAdjacentCoordinate: absoluteCoordinate.Build(3, 1),
		},
		{
			name:                  "Y is adjacent",
			adjacentCoordinate:    absoluteCoordinate.Build(1, 2),
			nonAdjacentCoordinate: absoluteCoordinate.Build(1, 3),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.True(t, aCoordinate.IsAdjacentTo(*testCase.adjacentCoordinate))
			assert.False(t, aCoordinate.IsAdjacentTo(*testCase.nonAdjacentCoordinate))
		})
	}
}
