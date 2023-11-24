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

func TestDoesNotProceedIfRepoErrors(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(curiosity, errors.New("repo error"))
	curiosity.On("Id").Return(uuid.New())

	err := moveUseCase.MoveSequenceAborting(curiosity.Id().String(), "f")

	curiosity.AssertNotCalled(t, "MoveForward")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Repository error")

	errs := moveUseCase.MoveSequence(curiosity.Id().String(), "f")

	curiosity.AssertNotCalled(t, "MoveForward")
	assert.Error(t, errs[0])
	assert.Contains(t, errs[0].Error(), "Repository error")
}

func TestDoesNotMoveANonExistingRover(t *testing.T) {
	repo := new(MockRepo)
	moveUseCase := move.For(repo)
	curiosity := new(MockRover)
	repo.On("GetRover").Return(nil, nil)

	curiosity.On("Id").Return(uuid.New())

	err := moveUseCase.MoveSequenceAborting(curiosity.Id().String(), "f")

	curiosity.AssertNotCalled(t, "MoveForward")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "no rover found")

	errs := moveUseCase.MoveSequence(curiosity.Id().String(), "f")

	curiosity.AssertNotCalled(t, "MoveForward")
	assert.Error(t, errs[0])
	assert.Contains(t, errs[0].Error(), "no rover found")
}
