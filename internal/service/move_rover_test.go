package service_test

import (
	"errors"
	"mars_rover/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandlesASingleMovementCommand(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("MoveForward").Return(nil)
	moveService := service.For(curiosity)

	movementErrors := moveService.MoveSequence("f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Nil(t, movementErrors)
}

func TestHandlesASingleTurningCommand(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("TurnRight").Return()
	moveService := service.For(curiosity)

	movementErrors := moveService.MoveSequence("r")

	curiosity.AssertCalled(t, "TurnRight")
	assert.Nil(t, movementErrors)
}

func TestHandlesASingleFailedMovementCommand(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("MoveForward").Return(errors.New(""))
	moveService := service.For(curiosity)

	movementErrors := moveService.MoveSequence("f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Error(t, movementErrors[0])
}

func TestHandlesASingleUnknownCommand(t *testing.T) {
	curiosity := new(MockRover)
	moveService := service.For(curiosity)

	movementErrors := moveService.MoveSequence("X")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertNotCalled(t, "TurnLeft")
	curiosity.AssertNotCalled(t, "MoveForward")
	curiosity.AssertNotCalled(t, "MoveBackward")
	assert.Error(t, movementErrors[0])
}

func TestHandlesMultipleKnownCommands(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("TurnRight").Return()
	curiosity.On("TurnLeft").Return()
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("MoveBackward").Return(nil)
	moveService := service.For(curiosity)

	movementErrors := moveService.MoveSequence("rlfb")

	curiosity.AssertCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Nil(t, movementErrors)
}

func TestHandlesMultipleErrors(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("MoveForward").Return(errors.New(""))
	curiosity.On("MoveBackward").Return(errors.New(""))
	moveService := service.For(curiosity)

	movementErrors := moveService.MoveSequence("fbXY")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertNotCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Len(t, movementErrors, 4)
}

func TestHandlesErrorsAndSuccessfulMovements(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("MoveBackward").Return(nil)
	curiosity.On("MoveForward").Return(errors.New(""))
	curiosity.On("TurnLeft").Return()
	moveService := service.For(curiosity)

	movementErrors := moveService.MoveSequence("bfXYl")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Len(t, movementErrors, 3)
}

type MockRover struct {
	mock.Mock
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
