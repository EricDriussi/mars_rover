package mocks

import (
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover/direction"
	. "mars_rover/src/domain/rover/planetMap"
	"mars_rover/src/domain/rover/id"
	. "mars_rover/src/domain/rover/id"
)

type MockRover struct {
	Mock
}

func (this *MockRover) TurnLeft() {
	this.Called()
}

func (this *MockRover) TurnRight() {
	this.Called()
}

func (this *MockRover) MoveForward() error {
	args := this.Called()
	return args.Error(0)
}

func (this *MockRover) MoveBackward() error {
	args := this.Called()
	return args.Error(0)
}

func (this *MockRover) Id() ID {
	args := this.Called()
	return args.Get(0).(ID)
}

func (this *MockRover) Coordinate() AbsoluteCoordinate {
	args := this.Called()
	return args.Get(0).(AbsoluteCoordinate)
}

func (this *MockRover) Direction() Direction {
	args := this.Called()
	return args.Get(0).(Direction)
}

func (this *MockRover) Map() Map {
	args := this.Called()
	return args.Get(0).(Map)
}

func RoverIn(planet Planet, coord AbsoluteCoordinate) *MockRover {
	mockRover := LandedRover(coord)
	mockRover.On("Map").Return(*OfPlanet(planet))
	return mockRover
}

func LandedRover(coord AbsoluteCoordinate) *MockRover {
	mockRover := new(MockRover)
	mockRover.On("Id").Return(id.New())
	mockRover.On("Direction").Return(North{})
	mockRover.On("Coordinate").Return(coord)
	return mockRover
}

func MakeAlwaysSuccessful(mockRover *MockRover) {
	mockRover.On("MoveForward").Return(nil)
	mockRover.On("MoveBackward").Return(nil)
	mockRover.On("TurnLeft").Return()
	mockRover.On("TurnRight").Return()
}
