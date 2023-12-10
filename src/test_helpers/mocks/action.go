package mocks

import (
	. "github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action"
	. "mars_rover/src/action/command"
	. "mars_rover/src/domain/rover"
)

type MockAction struct {
	Mock
}

func (this *MockAction) Random() (Rover, error) {
	args := this.Called()
	return args.Get(0).(Rover), args.Error(1)
}

func (this *MockAction) MoveSequence(roverId UUID, commands Commands) ([]MovementResult, error) {
	args := this.Called()
	return args.Get(0).([]MovementResult), args.Error(1)
}
