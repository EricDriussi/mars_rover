package location_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetsPosition(t *testing.T) {
	coord := *coordinate.NewAbsolute(1, 1)
	loc, _ := location.From(coord, &direction.North{})
	assert.Equal(t, coord, loc.Position())
}

func TestUpdatesPosition(t *testing.T) {
	loc, _ := location.From(*coordinate.NewAbsolute(1, 1), &direction.North{})
	loc.CalculatePositionAhead()
	loc.UpdatePosition()
	assert.Equal(t, *coordinate.NewAbsolute(1, 2), loc.Position())
	loc.CalculatePositionBehind()
	loc.UpdatePosition()
	assert.Equal(t, *coordinate.NewAbsolute(1, 1), loc.Position())
}

func TestResetsPosition(t *testing.T) {
	coord := *coordinate.NewAbsolute(1, 1)
	loc, _ := location.From(coord, &direction.North{})
	loc.CalculatePositionAhead()
	loc.Reset()
	assert.Equal(t, coord, loc.Position())
}

func TestFacesLeft(t *testing.T) {
	expectedDirection := &direction.West{}
	mockDirection := new(mockDirection)
	mockDirection.On("DirectionOnTheLeft").Return(expectedDirection)
	startingLocation, _ := location.From(*coordinate.NewAbsolute(1, 1), mockDirection)

	startingLocation.FaceLeft()

	expectedLocation, _ := location.From(*coordinate.NewAbsolute(1, 1), expectedDirection)
	assert.Equal(t, expectedLocation, startingLocation)
}

func TestFacesRight(t *testing.T) {
	expectedDirection := &direction.East{}
	mockDirection := new(mockDirection)
	mockDirection.On("DirectionOnTheRight").Return(expectedDirection)
	startingLocation, _ := location.From(*coordinate.NewAbsolute(1, 1), mockDirection)

	startingLocation.FaceRight()

	expectedLocation, _ := location.From(*coordinate.NewAbsolute(1, 1), expectedDirection)
	assert.Equal(t, expectedLocation, startingLocation)
}

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
			_, err := location.From(*testCase.coordinate, &mockDirection{})
			assert.Error(t, err)
		})
	}
}

type mockDirection struct {
	mock.Mock
}

func (this *mockDirection) Degree() int {
	args := this.Called()
	return args.Int(0)
}

func (this *mockDirection) DirectionOnTheLeft() direction.Direction {
	args := this.Called()
	return args.Get(0).(direction.Direction)
}

func (this *mockDirection) DirectionOnTheRight() direction.Direction {
	args := this.Called()
	return args.Get(0).(direction.Direction)
}

func (this *mockDirection) RelativePositionAhead() coordinate.RelativeCoordinate {
	args := this.Called()
	return args.Get(0).(coordinate.RelativeCoordinate)
}

func (this *mockDirection) RelativePositionBehind() coordinate.RelativeCoordinate {
	args := this.Called()
	return args.Get(0).(coordinate.RelativeCoordinate)
}
