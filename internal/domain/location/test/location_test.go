package location_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesNotAllowNegativeValues(t *testing.T) {
	testCases := []struct {
		name       string
		coordinate *coordinate.AbsoluteCoordinate
	}{
		{
			name:       "neither X nor Y can be negative",
			coordinate: coordinate.NewAbsolute(-1, -1),
		},
		{
			name:       "x cannot be negative",
			coordinate: coordinate.NewAbsolute(-1, 1),
		},
		{
			name:       "y cannot be negative",
			coordinate: coordinate.NewAbsolute(1, -1),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := location.From(*testCase.coordinate, mockDirection{})
			assert.Error(t, err)
		})
	}
}

type mockDirection struct{}

func (this mockDirection) Degree() int {
	return 420
}

func (this mockDirection) DirectionOnTheLeft() direction.Direction {
	return this
}

func (this mockDirection) DirectionOnTheRight() direction.Direction {
	return this
}

func (this mockDirection) RelativePositionAhead() coordinate.RelativeCoordinate {
	return *coordinate.RelativeFrom(0, 0)
}

func (this mockDirection) RelativePositionBehind() coordinate.RelativeCoordinate {
	return *coordinate.RelativeFrom(0, 0)
}
