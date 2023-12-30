package mocks

import (
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action/createRandom"
	. "mars_rover/src/action/move"
	. "mars_rover/src/action/move/command"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/uuid"
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

func (this *MockAction) Move(roverId UUID, commands Commands) ([]MovementResult, *MovementError) {
	args := this.Called()
	if args.Get(1) == nil {
		return args.Get(0).([]MovementResult), nil
	}
	return args.Get(0).([]MovementResult), args.Get(1).(*MovementError)
}
