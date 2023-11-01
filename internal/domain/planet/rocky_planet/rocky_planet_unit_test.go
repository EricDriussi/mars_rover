package rocky_planet_test

import (
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/test"
	rockyPlanet "mars_rover/internal/domain/planet/rocky_planet"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCanCreateIfNoMockObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.From(5, 5)
	obstacleOne, obstacleTwo := twoMockObstaclesWithinLimit(*sizeLimit)
	obstaclesWithinBounds := []obstacle.Obstacle{obstacleOne, obstacleTwo}
	_, err := rockyPlanet.Create(*sizeLimit, obstaclesWithinBounds)

	obstacleOne.AssertCalled(t, "IsBeyond", *sizeLimit)
	obstacleTwo.AssertCalled(t, "IsBeyond", *sizeLimit)
	assert.Nil(t, err)
}

func TestCannotCreateIfOneMockObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.From(5, 5)
	obstacleOne, obstacleTwo := twoMockObstaclesOneBeyondLimit(*sizeLimit)
	obstaclesWithinBounds := []obstacle.Obstacle{obstacleOne, obstacleTwo}
	_, err := rockyPlanet.Create(*sizeLimit, obstaclesWithinBounds)

	// asserting calls would depend on order of obstacle iteration
	assert.Error(t, err)
}

func twoMockObstaclesWithinLimit(limit size.Size) (*test.MockObstacle, *test.MockObstacle) {
	obstacleOne := new(test.MockObstacle)
	obstacleTwo := new(test.MockObstacle)
	obstacleOne.On("IsBeyond", mock.Anything).Return(false)
	obstacleTwo.On("IsBeyond", mock.Anything).Return(false)
	return obstacleOne, obstacleTwo
}

func twoMockObstaclesOneBeyondLimit(limit size.Size) (*test.MockObstacle, *test.MockObstacle) {
	obstacleOne := new(test.MockObstacle)
	obstacleTwo := new(test.MockObstacle)
	obstacleOne.On("IsBeyond", mock.Anything).Return(true)
	obstacleTwo.On("IsBeyond", mock.Anything).Return(false)
	return obstacleOne, obstacleTwo
}
