package bounded_random_creator_test

import (
	"errors"
	"mars_rover/src/action/createRandom/bounded"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoundedCreationDoesNotErrorIfRepoIsSuccessful(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(nil)
	repo.On("AddRover").Return(nil)
	act := bounded_random_creator.With(repo)

	for i := 0; i < 25; i++ {
		rover, err := act.Create()
		assert.Nil(t, err)
		assert.NotNil(t, rover)
	}
}

func TestBoundedCreationReportsRepoError(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(errors.New("repo error"))
	act := bounded_random_creator.With(repo)
	_, err := act.Create()
	assert.NotNil(t, err)
}

func TestBoundedCreationRespectsSensibleLimits(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(nil)
	repo.On("AddRover").Return(nil)
	act := bounded_random_creator.With(repo)

	for i := 0; i < 25; i++ {
		rover, err := act.Create()
		assert.Nil(t, err)

		planetMap := rover.Map()
		assert.GreaterOrEqual(t, planetMap.Width(), 4)
		assert.GreaterOrEqual(t, planetMap.Height(), 4)
		assert.LessOrEqual(t, planetMap.Width(), 20)
		assert.LessOrEqual(t, planetMap.Height(), 20)
		obstacles := planetMap.Obstacles()
		assert.GreaterOrEqual(t, len(obstacles.List()), 3)
		halfTheArea := planetMap.Width() * planetMap.Height() / 2
		assert.LessOrEqual(t, len(obstacles.List()), halfTheArea)
	}
}
