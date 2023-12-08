package action

import (
	. "github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action"
	. "mars_rover/src/domain/rover"
)

type MockRepo struct {
	Mock
}

func (this *MockRepo) Random() (Rover, error) {
	args := this.Called()
	return args.Get(0).(Rover), args.Error(1)
}

func (this *MockRepo) MoveSequence(roverId UUID, commands string) MovementResult {
	args := this.Called()
	return args.Get(0).(MovementResult)
}
