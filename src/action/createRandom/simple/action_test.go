package simple_random_creator_test

import (
	"errors"
	"mars_rover/src/action/createRandom/simple"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCreationDoesNotErrorIfRepoIsSuccessful(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(nil)
	repo.On("AddRover").Return(nil)
	act := simple_random_creator.With(repo)

	for i := 0; i < 25; i++ {
		rover, err := act.Create()
		assert.Nil(t, err)
		assert.NotNil(t, rover)
	}
}

func TestSimpleCreationReportsRepoError(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(errors.New("repo error"))
	act := simple_random_creator.With(repo)
	_, err := act.Create()
	assert.NotNil(t, err)
}
