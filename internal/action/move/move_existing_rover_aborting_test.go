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

func TestDoesNotAbortASingleSuccessfulMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	repo.On("UpdateRover").Return(nil)
	curiosity.On("MoveForward").Return(nil)

	err := moveUseCase.MoveSequenceAborting(curiosity.Id().String(), "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Nil(t, err)
}

func TestDoesNotAbortASingleRotationCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	repo.On("UpdateRover").Return(nil)
	curiosity.On("TurnRight").Return()

	err := moveUseCase.MoveSequenceAborting(curiosity.Id().String(), "r")

	curiosity.AssertCalled(t, "TurnRight")
	assert.Nil(t, err)
}

func TestAbortsASingleFailedMovementCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	repo.On("UpdateRover").Return(nil)
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(errors.New(movementBlockedError))

	err := moveUseCase.MoveSequenceAborting(curiosity.Id().String(), "f")

	curiosity.AssertCalled(t, "MoveForward")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), movementBlockedError)
}

func TestAbortsASingleUnknownCommand(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	repo.On("UpdateRover").Return(nil)

	err := moveUseCase.MoveSequenceAborting(curiosity.Id().String(), "X")

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
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	curiosity.On("TurnRight").Return()
	curiosity.On("TurnLeft").Return()
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("MoveBackward").Return(nil)
	repo.On("UpdateRover").Return(nil)

	err := moveUseCase.MoveSequenceAborting(curiosity.Id().String(), "rlfb")

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
	repo.On("GetRover").Return(curiosity, nil)
	curiosity.On("Id").Return(uuid.New())
	movementBlockedError := "movement blocked"
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("MoveBackward").Return(errors.New(movementBlockedError))
	curiosity.On("TurnRight").Return()
	repo.On("UpdateRover").Return(nil)

	err := moveUseCase.MoveSequenceAborting(curiosity.Id().String(), "fbr")

	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "MoveBackward")
	curiosity.AssertNotCalled(t, "TurnRight")
	assert.Contains(t, err.Error(), movementBlockedError)
}
