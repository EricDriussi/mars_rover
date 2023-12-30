package mocks

import (
	. "github.com/stretchr/testify/mock"
	"mars_rover/src/action/move/command"
	. "mars_rover/src/domain/rover"
)

type MockCommand struct {
	Mock
}

func (this *MockCommand) MapToRoverMovementFunction(rover Rover) command.RoverMovementFunc {
	args := this.Called(rover)
	return args.Get(0).(func() error)
}
