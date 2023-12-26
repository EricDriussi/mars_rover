package boundedRandomCreator_test

import (
	"errors"
	"mars_rover/src/action/createRandom/bounded"
	. "mars_rover/src/action/createRandom/bounded"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/rover/planetMap"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoundedCreationDoesNotErrorIfRepoIsSuccessful(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(nil)
	repo.On("AddRover").Return(nil)

	act := boundedRandomCreator.With(repo)
	rover, err := act.Create()

	assert.Nil(t, err)
	assert.NotNil(t, rover)
}

func TestBoundedCreationReportsRepoError(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(errors.New("repo error"))

	act := boundedRandomCreator.With(repo)
	_, err := act.Create()

	assert.Error(t, err)
}

func TestBoundedCreationRespectsSensibleLimits(t *testing.T) {
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(nil)
	repo.On("AddRover").Return(nil)
	act := boundedRandomCreator.With(repo)

	// since there is a lot of randomness involved, we create the game a bunch of times
	for i := 0; i < 25; i++ {
		rover, err := act.Create()
		assert.Nil(t, err)

		planetMap := rover.Map()
		assertPlanetMapIsWithin(t, planetMap)
		obstacles := planetMap.Obstacles()
		assertObstacleAmountIsWithinMinAndHalfTheArea(t, obstacles, planetMap)
	}
}

func assertObstacleAmountIsWithinMinAndHalfTheArea(t *testing.T, obstacles Obstacles, planetMap Map) {
	assert.GreaterOrEqual(t, len(obstacles.List()), MinObstacles)
	halfTheArea := planetMap.Width() * planetMap.Height() / 2
	assert.LessOrEqual(t, len(obstacles.List()), halfTheArea)
}

func assertPlanetMapIsWithin(t *testing.T, planetMap Map) {
	assert.GreaterOrEqual(t, planetMap.Width(), MinSize)
	assert.LessOrEqual(t, planetMap.Width(), MaxSize)
	assert.GreaterOrEqual(t, planetMap.Height(), MinSize)
	assert.LessOrEqual(t, planetMap.Height(), MaxSize)
}
