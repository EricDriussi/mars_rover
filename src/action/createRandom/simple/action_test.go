package simpleRandomCreator_test

import (
	"errors"
	simpleRandomCreator "mars_rover/src/action/createRandom/simple"
	"mars_rover/src/domain"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCreationDoesNotErrorIfRepoIsSuccessful(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(42, nil)
	repo.On("AddRover").Return(nil)

	act := simpleRandomCreator.With(repo, 10)
	rover, err := act.Create()

	assert.Nil(t, err)
	assert.NotNil(t, rover)
}

func TestSimpleCreationReportsRepoError(t *testing.T) {
	repo := new(MockRepo)
	repoErr := domain.PersistenceMalfunction(errors.New("repo error"))
	repo.On("AddPlanet").Return(0, repoErr)

	act := simpleRandomCreator.With(repo, 10)
	_, err := act.Create()

	assert.Error(t, err)
}
