package load_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/action/load"
	. "mars_rover/src/domain"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/rover/uuid"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestLoaderDoesNotErrorIfRepoReportsNoError(t *testing.T) {
	testRover := LandedRover(*absoluteCoordinate.Build(1, 1))
	repo := SuccessfulRepoFor(testRover)

	act := load.With(repo)
	loadedGame, err := act.Load(uuid.New())

	assert.Nil(t, err)
	assert.NotNil(t, loadedGame)
	assert.Equal(t, testRover, loadedGame.Rover)
}

func TestLoaderErrorsIfRepoReportsAnError(t *testing.T) {
	repo := new(MockRepo)
	repo.On("GetGame").Return(new(Game), NotFound())

	act := load.With(repo)
	_, err := act.Load(uuid.New())

	assert.Error(t, err)
	assert.True(t, err.IsNotFound())
}
