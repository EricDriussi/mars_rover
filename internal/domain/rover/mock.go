package rover

import (
	. "github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/rover/direction"
	. "mars_rover/internal/domain/rover/planetMap"
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

func (this *MockRover) Id() UUID {
	args := this.Called()
	return args.Get(0).(UUID)
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
