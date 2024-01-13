package gameLoader_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/action/gameLoader"
	. "mars_rover/src/domain"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/rover/id"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestLoaderDoesNotErrorIfRepoReportsNoError(t *testing.T) {
	testRover := LandedRover(*absoluteCoordinate.Build(1, 1))
	repo := SuccessfulRepoFor(testRover)

	act := gameLoader.With(repo)
	loadedGame, err := act.Load(id.New())

	assert.Nil(t, err)
	assert.NotNil(t, loadedGame)
	assert.Equal(t, testRover, loadedGame.Rover)
}

func TestLoaderErrorsIfRepoReportsAnError(t *testing.T) {
	repo := new(MockRepo)
	repo.On("GetGame").Return(new(Game), NotFound())

	act := gameLoader.With(repo)
	_, err := act.Load(id.New())

	assert.Error(t, err)
	assert.True(t, err.IsNotFound())
}
