package mocks

import (
	. "github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action"
	. "mars_rover/src/action/move"
	. "mars_rover/src/action/move/command"
	. "mars_rover/src/domain/rover"
)

type MockAction struct {
	Mock
}

func (this *MockAction) Create() (Rover, error) {
	args := this.Called()
	return args.Get(0).(Rover), args.Error(1)
}

func (this *MockAction) Move(roverId UUID, commands Commands) ([]MovementResult, *MovementError) {
	args := this.Called()
	if args.Get(1) == nil {
		return args.Get(0).([]MovementResult), nil
	}
	return args.Get(0).([]MovementResult), args.Get(1).(*MovementError)
}
