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
	rover := LandedRover(*absoluteCoordinate.Build(1, 1))
	repo := SuccessfulRepoFor(rover)
	loadAction := gameLoader.With(repo)

	loadedGame, err := loadAction.Load(id.New())

	assert.Nil(t, err)
	assert.NotNil(t, loadedGame)
	assert.Equal(t, rover, loadedGame.Rover)
}

func TestLoaderErrorsIfRepoReportsAnError(t *testing.T) {
	repo := new(MockRepo)
	repo.On("GetGame").Return(new(Game), NotFound())
	loadAction := gameLoader.With(repo)

	_, err := loadAction.Load(id.New())

	assert.Error(t, err)
	assert.True(t, err.IsNotFound())
}
