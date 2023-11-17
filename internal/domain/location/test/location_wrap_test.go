package location_test

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapsAheadIfOutOfBounds(t *testing.T) {
	maxValue := 2
	minValue := 0
	sizeLimit, _ := size.Square(maxValue)
	testCases := []struct {
		name      string
		starting  *absoluteCoordinate.AbsoluteCoordinate
		direction direction.Direction
		expected  *absoluteCoordinate.AbsoluteCoordinate
	}{
		{
			name:      "North",
			starting:  absoluteCoordinate.From(1, maxValue),
			direction: &direction.North{},
			expected:  absoluteCoordinate.From(1, minValue),
		},
		{
			name:      "South",
			starting:  absoluteCoordinate.From(1, minValue),
			direction: &direction.South{},
			expected:  absoluteCoordinate.From(1, maxValue),
		},
		{
			name:      "East",
			starting:  absoluteCoordinate.From(maxValue, 1),
			direction: &direction.East{},
			expected:  absoluteCoordinate.From(minValue, 1),
		},
		{
			name:      "West",
			starting:  absoluteCoordinate.From(minValue, 1),
			direction: &direction.West{},
			expected:  absoluteCoordinate.From(maxValue, 1),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			loc, _ := location.From(*testCase.starting, testCase.direction)
			loc.CalculatePositionAhead()
			loc.WrapAround(*sizeLimit)
			assert.Equal(t, *testCase.expected, loc.WillBeAt())
		})
	}
}

func TestDoesNotWrapAheadIfWithinBounds(t *testing.T) {
	loc, _ := location.From(*absoluteCoordinate.From(1, 1), &direction.North{})
	loc.CalculatePositionAhead()
	s, _ := size.Square(2)
	loc.WrapAround(*s)
	assert.Equal(t, *absoluteCoordinate.From(1, 2), loc.WillBeAt())
}

func TestWrapsBehindIfOutOfBounds(t *testing.T) {
	maxValue := 2
	minValue := 0
	sizeLimit, _ := size.Square(maxValue)
	testCases := []struct {
		name      string
		starting  *absoluteCoordinate.AbsoluteCoordinate
		direction direction.Direction
		expected  *absoluteCoordinate.AbsoluteCoordinate
	}{
		{
			name:      "North",
			starting:  absoluteCoordinate.From(1, minValue),
			direction: &direction.North{},
			expected:  absoluteCoordinate.From(1, maxValue),
		},
		{
			name:      "South",
			starting:  absoluteCoordinate.From(1, maxValue),
			direction: &direction.South{},
			expected:  absoluteCoordinate.From(1, minValue),
		},
		{
			name:      "East",
			starting:  absoluteCoordinate.From(minValue, 1),
			direction: &direction.East{},
			expected:  absoluteCoordinate.From(maxValue, 1),
		},
		{
			name:      "West",
			starting:  absoluteCoordinate.From(maxValue, 1),
			direction: &direction.West{},
			expected:  absoluteCoordinate.From(minValue, 1),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			loc, _ := location.From(*testCase.starting, testCase.direction)
			loc.CalculatePositionBehind()
			loc.WrapAround(*sizeLimit)
			assert.Equal(t, *testCase.expected, loc.WillBeAt())
		})
	}
}

func TestDoesNotWrapBehindIfWithinBounds(t *testing.T) {
	loc, _ := location.From(*absoluteCoordinate.From(1, 1), &direction.North{})
	loc.CalculatePositionBehind()
	s, _ := size.Square(2)
	loc.WrapAround(*s)
	assert.Equal(t, *absoluteCoordinate.From(1, 0), loc.WillBeAt())
}
