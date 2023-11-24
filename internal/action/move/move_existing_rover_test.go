package move_test

import (
	"errors"
	"github.com/google/uuid"
	"mars_rover/internal/action/move"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/infra/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlesASingleMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	repo.On("UpdateRover").Return(nil)
	curiosity.On("MoveForward").Return(nil)

	movementErrors := moveUseCase.MoveSequence(curiosity.Id().String(), "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Nil(t, movementErrors)
}

func TestHandlesASingleTurningCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	repo.On("UpdateRover").Return(nil)
	curiosity.On("TurnRight").Return()

	movementErrors := moveUseCase.MoveSequence(curiosity.Id().String(), "r")

	curiosity.AssertCalled(t, "TurnRight")
	assert.Nil(t, movementErrors)
}

func TestHandlesASingleFailedMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	repo.On("UpdateRover").Return(nil)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))

	movementErrors := moveUseCase.MoveSequence(curiosity.Id().String(), "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Error(t, movementErrors[0])
	assert.Contains(t, movementErrors[0].Error(), movementBlockedError)
}

func TestHandlesASingleUnknownCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	repo.On("UpdateRover").Return(nil)

	movementErrors := moveUseCase.MoveSequence(curiosity.Id().String(), "X")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertNotCalled(t, "TurnLeft")
	curiosity.AssertNotCalled(t, "MoveForward")
	curiosity.AssertNotCalled(t, "MoveBackward")
	assert.Error(t, movementErrors[0])
	assert.Contains(t, movementErrors[0].Error(), "invalid command")
}

func TestHandlesMultipleKnownCommands(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	curiosity.On("TurnRight").Return()
	curiosity.On("TurnLeft").Return()
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("MoveBackward").Return(nil)
	repo.On("UpdateRover").Return(nil)

	movementErrors := moveUseCase.MoveSequence(curiosity.Id().String(), "rlfb")

	curiosity.AssertCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Nil(t, movementErrors)
}

func TestHandlesMultipleErrors(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))
	curiosity.On("MoveBackward").Return(errors.New(movementBlockedError))
	repo.On("UpdateRover").Return(nil)

	movementErrors := moveUseCase.MoveSequence(curiosity.Id().String(), "fbXY")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertNotCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Len(t, movementErrors, 4)
	assert.Contains(t, movementErrors[0].Error(), movementBlockedError)
	assert.Contains(t, movementErrors[1].Error(), movementBlockedError)
	assert.Contains(t, movementErrors[2].Error(), "invalid command")
	assert.Contains(t, movementErrors[3].Error(), "invalid command")
}

func TestHandlesErrorsAndSuccessfulMovements(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	movementBlockedError := "movement blocked"
	curiosity.On("MoveBackward").Return(nil)
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))
	curiosity.On("TurnLeft").Return()
	repo.On("UpdateRover").Return(nil)

	movementErrors := moveUseCase.MoveSequence(curiosity.Id().String(), "bfXYl")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Len(t, movementErrors, 3)
	assert.Contains(t, movementErrors[0].Error(), movementBlockedError)
	assert.Contains(t, movementErrors[1].Error(), "invalid command")
	assert.Contains(t, movementErrors[2].Error(), "invalid command")
}
