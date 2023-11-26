package move_test

import (
	"errors"
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
	repo.On("UpdateRover").Return(nil)
	curiosity.On("MoveForward").Return(nil)

	_, movementErrors := moveUseCase.MoveSequence(curiosity, "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Nil(t, movementErrors)
}

func TestHandlesASingleTurningCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)
	curiosity.On("TurnRight").Return()

	_, movementErrors := moveUseCase.MoveSequence(curiosity, "r")

	curiosity.AssertCalled(t, "TurnRight")
	assert.Nil(t, movementErrors)
}

func TestHandlesASingleFailedMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))

	_, movementErrors := moveUseCase.MoveSequence(curiosity, "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Len(t, movementErrors, 1)
	assert.Error(t, movementErrors[0])
	assert.Contains(t, movementErrors[0].Error(), movementBlockedError)
}

func TestHandlesASingleUnknownCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)

	_, movementErrors := moveUseCase.MoveSequence(curiosity, "X")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertNotCalled(t, "TurnLeft")
	curiosity.AssertNotCalled(t, "MoveForward")
	curiosity.AssertNotCalled(t, "MoveBackward")
	assert.Len(t, movementErrors, 1)
	assert.Error(t, movementErrors[0])
	assert.Contains(t, movementErrors[0].Error(), "invalid command")
}

func TestHandlesMultipleKnownCommands(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	curiosity.On("TurnRight").Return()
	curiosity.On("TurnLeft").Return()
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("MoveBackward").Return(nil)
	repo.On("UpdateRover").Return(nil)

	_, movementErrors := moveUseCase.MoveSequence(curiosity, "rlfb")

	curiosity.AssertCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Nil(t, movementErrors)
}

func TestHandlesMultipleMovementErrors(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))
	curiosity.On("MoveBackward").Return(errors.New(movementBlockedError))
	repo.On("UpdateRover").Return(nil)

	_, movementErrors := moveUseCase.MoveSequence(curiosity, "fbXY")

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
	movementBlockedError := "movement blocked"
	curiosity.On("MoveBackward").Return(nil)
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))
	curiosity.On("TurnLeft").Return()
	repo.On("UpdateRover").Return(nil)

	_, movementErrors := moveUseCase.MoveSequence(curiosity, "bfXYl")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Len(t, movementErrors, 3)
	assert.Contains(t, movementErrors[0].Error(), movementBlockedError)
	assert.Contains(t, movementErrors[1].Error(), "invalid command")
	assert.Contains(t, movementErrors[2].Error(), "invalid command")
}

func TestReportsRepoError(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	curiosity.On("MoveForward").Return(nil)
	repoError := "repo error"
	repo.On("UpdateRover").Return(errors.New(repoError))

	_, errs := moveUseCase.MoveSequence(curiosity, "f")

	assert.Len(t, errs, 1)
	assert.Error(t, errs[0])
	assert.Contains(t, errs[0].Error(), repoError)
}

func TestHandlesNilRover(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)

	_, errs := moveUseCase.MoveSequence(nil, "f")

	assert.Len(t, errs, 1)
	assert.Error(t, errs[0])
	assert.Contains(t, errs[0].Error(), "got nil rover")
}
