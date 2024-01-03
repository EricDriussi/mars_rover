package planetWithObstacles_test

import (
	. "mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/planet/planetWithObstacles"
	"mars_rover/src/domain/size"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
)

func TestCanCreateIfNoMockObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstacleOne := new(MockObstacle)
	obstacleTwo := new(MockObstacle)
	obstacleOne.On("IsBeyond", Anything).Return(false)
	obstacleTwo.On("IsBeyond", Anything).Return(false)
	obstaclesWithinBounds := []Obstacle{obstacleOne, obstacleTwo}
	_, err := planetWithObstacles.Create("testColor", *sizeLimit, obstaclesWithinBounds)

	obstacleOne.AssertCalled(t, "IsBeyond", *sizeLimit)
	obstacleTwo.AssertCalled(t, "IsBeyond", *sizeLimit)
	assert.Nil(t, err)
}

func TestCannotCreateIfOneMockObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstacleOne := new(MockObstacle)
	obstacleTwo := new(MockObstacle)
	obstacleOne.On("IsBeyond", Anything).Return(true)
	obstacleTwo.On("IsBeyond", Anything).Return(false)
	obstaclesWithinBounds := []Obstacle{obstacleOne, obstacleTwo}
	_, err := planetWithObstacles.Create("testColor", *sizeLimit, obstaclesWithinBounds)

	assert.Error(t, err)
}

func TestCannotCreateIfSizeTooSmall(t *testing.T) {
	sizeLimit, _ := size.Square(1)
	_, err := planetWithObstacles.Create("testColor", *sizeLimit, []Obstacle{})

	assert.Error(t, err)
}

func TestCannotCreateIfNoObstacles(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	_, err := planetWithObstacles.Create("testColor", *sizeLimit, []Obstacle{})

	assert.Error(t, err)
}