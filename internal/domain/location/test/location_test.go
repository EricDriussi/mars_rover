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
	testLocation, _ := location.From(coord, &direction.North{})
	assert.Equal(t, coord, testLocation.Position())
}

func TestGetsDirection(t *testing.T) {
	direction := &direction.North{}
	testLocation, _ := location.From(*coordinate.NewAbsolute(1, 1), direction)
	assert.Equal(t, direction, testLocation.Direction())
}

func TestUpdatesPosition(t *testing.T) {
	testLocation, _ := location.From(*coordinate.NewAbsolute(1, 1), &direction.North{})
	testLocation.CalculatePositionAhead()
	testLocation.UpdatePosition()
	assert.Equal(t, *coordinate.NewAbsolute(1, 2), testLocation.Position())
	testLocation.CalculatePositionBehind()
	testLocation.UpdatePosition()
	assert.Equal(t, *coordinate.NewAbsolute(1, 1), testLocation.Position())
}

func TestResetsPosition(t *testing.T) {
	coord := *coordinate.NewAbsolute(1, 1)
	testLocation, _ := location.From(coord, &direction.North{})
	testLocation.CalculatePositionAhead()
	testLocation.Reset()
	assert.Equal(t, coord, testLocation.Position())
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
