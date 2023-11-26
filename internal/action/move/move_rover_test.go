package move_test

import (
	"errors"
	"mars_rover/internal/action/move"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/infra/test"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlesASingleMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)
	curiosity.On("MoveForward").Return(nil)

	movementResult := moveUseCase.MoveSequence(curiosity, "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Nil(t, movementResult.Error)
	assert.False(t, movementResult.HasMovementErrors())
}

func TestHandlesASingleTurningCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)
	curiosity.On("TurnRight").Return()

	movementResult := moveUseCase.MoveSequence(curiosity, "r")

	curiosity.AssertCalled(t, "TurnRight")
	assert.Nil(t, movementResult.Error)
	assert.False(t, movementResult.HasMovementErrors())
}

func TestHandlesASingleFailedMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))

	movementResult := moveUseCase.MoveSequence(curiosity, "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Len(t, movementResult.MovementErrors.List(), 1)
	assert.True(t, contains(movementResult.MovementErrors.AsStringArray(), movementBlockedError))
}

func TestHandlesASingleUnknownCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)

	movementResult := moveUseCase.MoveSequence(curiosity, "X")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertNotCalled(t, "TurnLeft")
	curiosity.AssertNotCalled(t, "MoveForward")
	curiosity.AssertNotCalled(t, "MoveBackward")
	assert.Len(t, movementResult.MovementErrors.List(), 1)
	assert.True(t, contains(movementResult.MovementErrors.AsStringArray(), "invalid command"))
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

	movementResult := moveUseCase.MoveSequence(curiosity, "rlfb")

	curiosity.AssertCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Nil(t, movementResult.Error)
	assert.Nil(t, movementResult.MovementErrors.List())
}

func TestHandlesMultipleMovementErrors(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))
	curiosity.On("MoveBackward").Return(errors.New(movementBlockedError))
	repo.On("UpdateRover").Return(nil)

	movementResult := moveUseCase.MoveSequence(curiosity, "fbXY")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertNotCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Len(t, movementResult.MovementErrors.List(), 4)
	assert.Contains(t, movementResult.MovementErrors.AsStringArray()[0], movementBlockedError)
	assert.Contains(t, movementResult.MovementErrors.AsStringArray()[1], movementBlockedError)
	assert.Contains(t, movementResult.MovementErrors.AsStringArray()[2], "invalid command")
	assert.Contains(t, movementResult.MovementErrors.AsStringArray()[3], "invalid command")
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

	movementResult := moveUseCase.MoveSequence(curiosity, "bfXYl")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Len(t, movementResult.MovementErrors.List(), 3)
	assert.Contains(t, movementResult.MovementErrors.AsStringArray()[0], movementBlockedError)
	assert.Contains(t, movementResult.MovementErrors.AsStringArray()[1], "invalid command")
	assert.Contains(t, movementResult.MovementErrors.AsStringArray()[2], "invalid command")
}

func TestReportsRepoError(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	curiosity.On("MoveForward").Return(nil)
	repoError := "repo error"
	repo.On("UpdateRover").Return(errors.New(repoError))

	movementResult := moveUseCase.MoveSequence(curiosity, "f")

	assert.Nil(t, movementResult.MovementErrors)
	assert.NotNil(t, movementResult.Error)
	assert.Contains(t, movementResult.Error.Error(), repoError)
}

func TestHandlesNilRover(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)

	movementResult := moveUseCase.MoveSequence(nil, "f")

	assert.Nil(t, movementResult.MovementErrors)
	assert.NotNil(t, movementResult.Error)
	assert.Contains(t, movementResult.Error.Error(), "got nil rover")
}

func contains(stringArray []string, str string) bool {
	for _, s := range stringArray {
		if strings.Contains(s, str) {
			return true
		}
	}
	return false
}
