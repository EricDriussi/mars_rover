package location_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: mocks?
func TestWrapsAroundIfOutOfBounds(t *testing.T) {
	maxValue := 2
	minValue := 0
	sizeLimit := size.Size{Width: maxValue, Height: maxValue}
	testCases := []struct {
		name      string
		starting  *coordinate.AbsoluteCoordinate
		direction direction.Direction
		expected  *coordinate.AbsoluteCoordinate
	}{
		{
			name:      "North",
			starting:  coordinate.NewAbsolute(1, maxValue),
			direction: &direction.North{},
			expected:  coordinate.NewAbsolute(1, minValue),
		},
		{
			name:      "South",
			starting:  coordinate.NewAbsolute(1, minValue),
			direction: &direction.South{},
			expected:  coordinate.NewAbsolute(1, maxValue),
		},
		{
			name:      "East",
			starting:  coordinate.NewAbsolute(maxValue, 1),
			direction: &direction.East{},
			expected:  coordinate.NewAbsolute(minValue, 1),
		},
		{
			name:      "West",
			starting:  coordinate.NewAbsolute(minValue, 1),
			direction: &direction.West{},
			expected:  coordinate.NewAbsolute(maxValue, 1),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			loc, _ := location.From(*testCase.starting, testCase.direction)
			loc.StartMovingAhead()
			loc.WrapAround(sizeLimit)
			assert.Equal(t, *testCase.expected, loc.WillBeAt())
		})
	}
}
