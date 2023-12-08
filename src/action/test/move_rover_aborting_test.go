package action_test

import (
	"errors"
	"mars_rover/src/action"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesNotAbortASingleSuccessfulMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	act := action.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)
	curiosity.On("MoveForward").Return(nil)

	_, err := act.MoveSequenceAborting(curiosity, "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Nil(t, err)
}

func TestDoesNotAbortASingleRotationCommand(t *testing.T) {
	repo := new(MockRepo)
	act := action.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)
	curiosity.On("TurnRight").Return()

	_, err := act.MoveSequenceAborting(curiosity, "r")

	curiosity.AssertCalled(t, "TurnRight")
	assert.Nil(t, err)
}

func TestAbortsASingleFailedMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	act := action.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))

	_, err := act.MoveSequenceAborting(curiosity, "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), movementBlockedError)
}

func TestAbortsASingleUnknownCommand(t *testing.T) {
	repo := new(MockRepo)
	act := action.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)

	_, err := act.MoveSequenceAborting(curiosity, "X")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertNotCalled(t, "TurnLeft")
	curiosity.AssertNotCalled(t, "MoveForward")
	curiosity.AssertNotCalled(t, "MoveBackward")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid command")
}

func TestDoesNotAbortMultipleKnownCommands(t *testing.T) {
	repo := new(MockRepo)
	act := action.For(repo)
	curiosity := new(MockRover)
	curiosity.On("TurnRight").Return()
	curiosity.On("TurnLeft").Return()
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("MoveBackward").Return(nil)
	repo.On("UpdateRover").Return(nil)

	_, err := act.MoveSequenceAborting(curiosity, "rlfb")

	curiosity.AssertCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Nil(t, err)
}

func TestAbortsOnFirstFailure(t *testing.T) {
	repo := new(MockRepo)
	act := action.For(repo)
	curiosity := new(MockRover)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("MoveBackward").Return(errors.New(movementBlockedError))
	curiosity.On("TurnRight").Return()
	repo.On("UpdateRover").Return(nil)

	_, err := act.MoveSequenceAborting(curiosity, "fbr")

	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	curiosity.AssertNotCalled(t, "TurnRight")
	assert.Contains(t, err.Error(), movementBlockedError)
}

func TestAbortsOnRepoError(t *testing.T) {
	repo := new(MockRepo)
	act := action.For(repo)
	curiosity := new(MockRover)
	curiosity.On("MoveForward").Return(nil)
	repoError := "repo error"
	repo.On("UpdateRover").Return(errors.New(repoError))

	_, err := act.MoveSequenceAborting(curiosity, "f")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), repoError)
}

func TestAbortsOnNilRover(t *testing.T) {
	repo := new(MockRepo)
	act := action.For(repo)

	_, err := act.MoveSequenceAborting(nil, "f")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "got nil rover")
}
