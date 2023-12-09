package action_test

import (
	"errors"
	"github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	"mars_rover/src/action"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence/test"
	. "mars_rover/src/test_helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlesASingleMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	curiosity.On("MoveForward").Return(nil)
	repo.On("UpdateRover").Return(nil)

	movementResult := act.MoveSequence(uuid.New(), "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Nil(t, movementResult.Error)
	assertHasNoMovementErrors(t, movementResult)
}

func TestHandlesASingleTurningCommand(t *testing.T) {
	repo := new(MockRepo)
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	curiosity.On("TurnRight").Return()
	repo.On("UpdateRover").Return(nil)

	movementResult := act.MoveSequence(uuid.New(), "r")

	curiosity.AssertCalled(t, "TurnRight")
	assert.Nil(t, movementResult.Error)
	assertHasNoMovementErrors(t, movementResult)
}

func TestHandlesASingleFailedMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))
	repo.On("UpdateRover").Return(nil)

	movementResult := act.MoveSequence(uuid.New(), "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Len(t, movementResult.MovementErrors.List(), 1)
	AssertContains(t, movementResult.MovementErrors.AsStringArray(), movementBlockedError)
}

func TestHandlesASingleUnknownCommand(t *testing.T) {
	repo := new(MockRepo)
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(nil)

	movementResult := act.MoveSequence(uuid.New(), "X")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertNotCalled(t, "TurnLeft")
	curiosity.AssertNotCalled(t, "MoveForward")
	curiosity.AssertNotCalled(t, "MoveBackward")
	assert.Len(t, movementResult.MovementErrors.List(), 1)
	AssertContains(t, movementResult.MovementErrors.AsStringArray(), "invalid command")
}

func TestHandlesMultipleKnownCommands(t *testing.T) {
	repo := new(MockRepo)
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	curiosity.On("TurnRight").Return()
	curiosity.On("TurnLeft").Return()
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("MoveBackward").Return(nil)
	repo.On("UpdateRover").Return(nil)

	movementResult := act.MoveSequence(uuid.New(), "rlfb")

	curiosity.AssertCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Nil(t, movementResult.Error)
	assert.Nil(t, movementResult.MovementErrors.List())
}

func TestHandlesMultipleMovementErrors(t *testing.T) {
	repo := new(MockRepo)
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))
	curiosity.On("MoveBackward").Return(errors.New(movementBlockedError))
	repo.On("UpdateRover").Return(nil)

	movementResult := act.MoveSequence(uuid.New(), "fbXY")

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
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveBackward").Return(nil)
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))
	curiosity.On("TurnLeft").Return()
	repo.On("UpdateRover").Return(nil)

	movementResult := act.MoveSequence(uuid.New(), "bfXYl")

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
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	curiosity.On("MoveForward").Return(nil)
	repoError := "repo error"
	repo.On("UpdateRover").Return(errors.New(repoError))

	movementResult := act.MoveSequence(uuid.New(), "f")

	assert.Nil(t, movementResult.MovementErrors)
	assert.NotNil(t, movementResult.Error)
	assert.Contains(t, movementResult.Error.Error(), repoError)
}

func TestHandlesNonExistingRover(t *testing.T) {
	repo := new(MockRepo)
	curiosity := new(MockRover)
	act := action.For(repo)
	repoError := "got nil rover"
	repo.On("GetRover", Anything).Return(curiosity, errors.New(repoError))

	movementResult := act.MoveSequence(uuid.New(), "f")

	assert.Nil(t, movementResult.MovementErrors)
	assert.NotNil(t, movementResult.Error)
	assert.Contains(t, movementResult.Error.Error(), repoError)
}

func assertHasNoMovementErrors(t *testing.T, result action.MovementResult) {
	hasErrors := result.MovementErrors != nil && len(result.MovementErrors.List()) > 0
	assert.False(t, hasErrors)
}
