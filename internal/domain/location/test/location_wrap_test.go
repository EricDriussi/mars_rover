package location_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapsAheadIfOutOfBounds(t *testing.T) {
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
			loc.CalculatePositionAhead()
			loc.WrapAround(sizeLimit)
			assert.Equal(t, *testCase.expected, loc.WillBeAt())
		})
	}
}

func TestDoesNotWrapAheadIfWithinBounds(t *testing.T) {
	loc, _ := location.From(*coordinate.NewAbsolute(1, 1), &direction.North{})
	loc.CalculatePositionAhead()
	loc.WrapAround(size.Size{Width: 2, Height: 2})
	assert.Equal(t, *coordinate.NewAbsolute(1, 2), loc.WillBeAt())
}

func TestWrapsBehindIfOutOfBounds(t *testing.T) {
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
			starting:  coordinate.NewAbsolute(1, minValue),
			direction: &direction.North{},
			expected:  coordinate.NewAbsolute(1, maxValue),
		},
		{
			name:      "South",
			starting:  coordinate.NewAbsolute(1, maxValue),
			direction: &direction.South{},
			expected:  coordinate.NewAbsolute(1, minValue),
		},
		{
			name:      "East",
			starting:  coordinate.NewAbsolute(minValue, 1),
			direction: &direction.East{},
			expected:  coordinate.NewAbsolute(maxValue, 1),
		},
		{
			name:      "West",
			starting:  coordinate.NewAbsolute(maxValue, 1),
			direction: &direction.West{},
			expected:  coordinate.NewAbsolute(minValue, 1),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			loc, _ := location.From(*testCase.starting, testCase.direction)
			loc.CalculatePositionBehind()
			loc.WrapAround(sizeLimit)
			assert.Equal(t, *testCase.expected, loc.WillBeAt())
		})
	}
}

func TestDoesNotWrapBehindIfWithinBounds(t *testing.T) {
	loc, _ := location.From(*coordinate.NewAbsolute(1, 1), &direction.North{})
	loc.CalculatePositionBehind()
	loc.WrapAround(size.Size{Width: 2, Height: 2})
	assert.Equal(t, *coordinate.NewAbsolute(1, 0), loc.WillBeAt())
}
