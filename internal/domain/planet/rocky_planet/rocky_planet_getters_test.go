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

func TestGetsSize(t *testing.T) {
	sizeLimit, _ := size.From(5, 5)
	planet, _ := rockyPlanet.Create(*sizeLimit, []obstacle.Obstacle{})
	assert.Equal(t, *sizeLimit, planet.Size())
}

func TestGetsObstacles(t *testing.T) {
	sizeLimit, _ := size.From(5, 5)
	obstacleOne := new(test.MockObstacle)
	obstacleTwo := new(test.MockObstacle)
	obstacles := []obstacle.Obstacle{obstacleOne, obstacleTwo}
	obstacleOne.On("IsBeyond", mock.Anything).Return(false)
	obstacleTwo.On("IsBeyond", mock.Anything).Return(false)
	planet, _ := rockyPlanet.Create(*sizeLimit, obstacles)
	assert.Equal(t, obstacles, planet.Obstacles())
}
