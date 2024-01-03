package planetWithObstacles_test

import (
	. "mars_rover/src/domain/obstacle"
	obs "mars_rover/src/domain/obstacle/obstacles"
	"mars_rover/src/domain/planet/planetWithObstacles"
	"mars_rover/src/domain/size"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
)

func TestGetsSize(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstacle := new(MockObstacle)
	obstacle.On("IsBeyond", Anything).Return(false)
	planet, _ := planetWithObstacles.Create("testColor", *sizeLimit, []Obstacle{obstacle})

	assert.Equal(t, *sizeLimit, planet.Size())
}

func TestGetsObstacles(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstacleOne := new(MockObstacle)
	obstacleTwo := new(MockObstacle)
	obstacles := []Obstacle{obstacleOne, obstacleTwo}
	obstacleOne.On("IsBeyond", Anything).Return(false)
	obstacleTwo.On("IsBeyond", Anything).Return(false)
	planet, _ := planetWithObstacles.Create("testColor", *sizeLimit, obstacles)

	assert.Equal(t, *obs.FromList(obstacles), planet.Obstacles())
}

func TestGetsColor(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstacleOne := new(MockObstacle)
	obstacleTwo := new(MockObstacle)
	obstacles := []Obstacle{obstacleOne, obstacleTwo}
	obstacleOne.On("IsBeyond", Anything).Return(false)
	obstacleTwo.On("IsBeyond", Anything).Return(false)
	color := "aColor"
	planet, _ := planetWithObstacles.Create(color, *sizeLimit, obstacles)

	assert.Equal(t, color, planet.Color())
}