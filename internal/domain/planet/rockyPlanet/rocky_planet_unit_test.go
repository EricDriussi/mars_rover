package rockyPlanet_test

import (
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/test"
	"mars_rover/internal/domain/planet/rockyPlanet"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCanCreateIfNoMockObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstacleOne := new(test.MockObstacle)
	obstacleTwo := new(test.MockObstacle)
	obstacleOne.On("IsBeyond", mock.Anything).Return(false)
	obstacleTwo.On("IsBeyond", mock.Anything).Return(false)
	obstaclesWithinBounds := []obstacle.Obstacle{obstacleOne, obstacleTwo}
	_, err := rockyPlanet.Create("testColor", *sizeLimit, obstaclesWithinBounds)

	obstacleOne.AssertCalled(t, "IsBeyond", *sizeLimit)
	obstacleTwo.AssertCalled(t, "IsBeyond", *sizeLimit)
	assert.Nil(t, err)
}

func TestCannotCreateIfOneMockObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstacleOne := new(test.MockObstacle)
	obstacleTwo := new(test.MockObstacle)
	obstacleOne.On("IsBeyond", mock.Anything).Return(true)
	obstacleTwo.On("IsBeyond", mock.Anything).Return(false)
	obstaclesWithinBounds := []obstacle.Obstacle{obstacleOne, obstacleTwo}
	_, err := rockyPlanet.Create("testColor", *sizeLimit, obstaclesWithinBounds)

	// asserting calls would depend on order of obstacle iteration
	assert.Error(t, err)
}
