package action_test

import (
	"errors"
	"github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	"mars_rover/src/action"
	"mars_rover/src/action/command"
	. "mars_rover/src/test_helpers"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlesASingleSuccessfulMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	curiosity.On("MoveForward").Return(nil)
	repo.On("UpdateRover").Return(nil)
	commands := command.Commands{command.Forward}

	movementResult, err := act.MoveSequence(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "MoveForward")
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
	commands := command.Commands{command.Forward}

	movementResult, err := act.MoveSequence(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "MoveForward")
	assert.Len(t, movementResult.Collisions.List(), 1)
	AssertContains(t, movementResult.Collisions.AsStringArray(), movementBlockedError)
}

func TestHandlesASingleTurningCommand(t *testing.T) {
	repo := new(MockRepo)
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	curiosity.On("TurnRight").Return()
	repo.On("UpdateRover").Return(nil)
	commands := command.Commands{command.Right}

	movementResult, err := act.MoveSequence(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "TurnRight")
	assertHasNoMovementErrors(t, movementResult)
}

func TestHandlesACombinationOfSuccessfulAndUnsuccessfulCommands(t *testing.T) {
	repo := new(MockRepo)
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveBackward").Return(nil)
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))
	curiosity.On("TurnLeft").Return()
	repo.On("UpdateRover").Return(nil)
	commands := command.Commands{command.Backward, command.Forward, command.Left}

	movementResult, err := act.MoveSequence(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Len(t, movementResult.Collisions.List(), 1)
	assert.Contains(t, movementResult.Collisions.AsStringArray()[0], movementBlockedError)
}

func TestReportsRepoErrorWhenGettingRover(t *testing.T) {
	repo := new(MockRepo)
	act := action.For(repo)
	repoError := "repo error"
	repo.On("GetRover", Anything).Return(new(MockRover), errors.New(repoError))
	irrelevantCommand := command.Forward

	movementResult, err := act.MoveSequence(uuid.New(), command.Commands{irrelevantCommand})

	assert.Nil(t, movementResult.Collisions)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), repoError)
}

func TestReportsRepoErrorWhenUpdatingRover(t *testing.T) {
	repo := new(MockRepo)
	curiosity := new(MockRover)
	act := action.For(repo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	curiosity.On("MoveForward").Return(nil)
	repoError := "repo error"
	repo.On("UpdateRover").Return(errors.New(repoError))
	irrelevantCommand := command.Forward

	movementResult, err := act.MoveSequence(uuid.New(), command.Commands{irrelevantCommand})

	assert.Nil(t, movementResult.Collisions)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), repoError)
}

func assertHasNoMovementErrors(t *testing.T, result action.MovementResult) {
	hasErrors := result.Collisions != nil && len(result.Collisions.List()) > 0
	assert.False(t, hasErrors)
}
