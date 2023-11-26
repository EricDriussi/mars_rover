package move_test

import (
	"errors"
	"mars_rover/internal/action/move"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/infra/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesNotAbortASingleSuccessfulMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)
	curiosity.On("MoveForward").Return(nil)

	_, err := moveUseCase.MoveSequenceAborting(curiosity, "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Nil(t, err)
}

func TestDoesNotAbortASingleRotationCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)
	curiosity.On("TurnRight").Return()

	_, err := moveUseCase.MoveSequenceAborting(curiosity, "r")

	curiosity.AssertCalled(t, "TurnRight")
	assert.Nil(t, err)
}

func TestAbortsASingleFailedMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))

	_, err := moveUseCase.MoveSequenceAborting(curiosity, "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), movementBlockedError)
}

func TestAbortsASingleUnknownCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("UpdateRover").Return(nil)

	_, err := moveUseCase.MoveSequenceAborting(curiosity, "X")

	curiosity.AssertNotCalled(t, "TurnRight")
	curiosity.AssertNotCalled(t, "TurnLeft")
	curiosity.AssertNotCalled(t, "MoveForward")
	curiosity.AssertNotCalled(t, "MoveBackward")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid command")
}

func TestDoesNotAbortMultipleKnownCommands(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	curiosity.On("TurnRight").Return()
	curiosity.On("TurnLeft").Return()
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("MoveBackward").Return(nil)
	repo.On("UpdateRover").Return(nil)

	_, err := moveUseCase.MoveSequenceAborting(curiosity, "rlfb")

	curiosity.AssertCalled(t, "TurnRight")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	assert.Nil(t, err)
}

func TestAbortsOnFirstFailure(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("MoveBackward").Return(errors.New(movementBlockedError))
	curiosity.On("TurnRight").Return()
	repo.On("UpdateRover").Return(nil)

	_, err := moveUseCase.MoveSequenceAborting(curiosity, "fbr")

	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	curiosity.AssertNotCalled(t, "TurnRight")
	assert.Contains(t, err.Error(), movementBlockedError)
}

func TestAbortsOnRepoError(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	curiosity.On("MoveForward").Return(nil)
	repoError := "repo error"
	repo.On("UpdateRover").Return(errors.New(repoError))

	_, err := moveUseCase.MoveSequenceAborting(curiosity, "f")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), repoError)
}

func TestAbortsOnNilRover(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)

	_, err := moveUseCase.MoveSequenceAborting(nil, "f")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "got nil rover")
}
