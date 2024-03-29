package simpleRandomCreator_test

import (
	"errors"
	"mars_rover/src/action/randomCreator/simple"
	. "mars_rover/src/domain"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCreationDoesNotErrorIfRepoIsSuccessful(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(42, nil)
	repo.On("AddRover").Return(nil)
	createAction := simpleRandomCreator.With(repo, 10)

	rover, err := createAction.Create()

	assert.Nil(t, err)
	assert.NotNil(t, rover)
}

func TestSimpleCreationReportsRepoError(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(-1, PersistenceMalfunction(errors.New("repo error")))
	createAction := simpleRandomCreator.With(repo, 10)

	_, err := createAction.Create()

	assert.Error(t, err)
}
