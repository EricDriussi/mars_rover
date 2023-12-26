package simpleRandomCreator_test

import (
	"errors"
	simpleRandomCreator "mars_rover/src/action/createRandom/simple"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCreationDoesNotErrorIfRepoIsSuccessful(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(nil)
	repo.On("AddRover").Return(nil)

	act := simpleRandomCreator.With(repo)
	rover, err := act.Create()

	assert.Nil(t, err)
	assert.NotNil(t, rover)
}

func TestSimpleCreationReportsRepoError(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(errors.New("repo error"))

	act := simpleRandomCreator.With(repo)
	_, err := act.Create()

	assert.Error(t, err)
}
