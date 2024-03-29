package mocks

import (
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action/gameLoader"
	. "mars_rover/src/action/mover"
	. "mars_rover/src/action/mover/command"
	. "mars_rover/src/action/randomCreator"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/id"
)

type MockAction struct {
	Mock
}

func (this *MockAction) Create() (Rover, *CreationError) {
	args := this.Called()
	if args.Get(1) == nil {
		return args.Get(0).(Rover), nil
	}
	return args.Get(0).(Rover), args.Get(1).(*CreationError)
}

func (this *MockAction) Move(ID, Commands) ([]MovementResult, *MovementError) {
	args := this.Called()
	if args.Get(1) == nil {
		return args.Get(0).([]MovementResult), nil
	}
	return args.Get(0).([]MovementResult), args.Get(1).(*MovementError)
}

func (this *MockAction) Load(ID) (*Game, *LoadError) {
	args := this.Called()
	if args.Get(1) == nil {
		return args.Get(0).(*Game), nil
	}
	return args.Get(0).(*Game), args.Get(1).(*LoadError)
}
