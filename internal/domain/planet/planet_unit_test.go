package planet_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCanCreateIfNoMockObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.From(5, 5)
	obstacleOne, obstacleTwo := twoMockObstaclesWithinLimit(*sizeLimit)
	obstaclesWithinBounds := []obstacle.Obstacle{obstacleOne, obstacleTwo}
	_, err := planet.Create(*sizeLimit, obstaclesWithinBounds)

	obstacleOne.AssertCalled(t, "IsBeyond", *sizeLimit)
	obstacleTwo.AssertCalled(t, "IsBeyond", *sizeLimit)
	assert.Nil(t, err)
}

func TestCannotCreateIfOneMockObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.From(5, 5)
	obstacleOne, obstacleTwo := twoMockObstaclesOneBeyondLimit(*sizeLimit)
	obstaclesWithinBounds := []obstacle.Obstacle{obstacleOne, obstacleTwo}
	_, err := planet.Create(*sizeLimit, obstaclesWithinBounds)

	// asserting calls would depend on order of obstacle iteration
	assert.Error(t, err)
}

func twoMockObstaclesWithinLimit(limit size.Size) (*MockObstacle, *MockObstacle) {
	obstacleOne := new(MockObstacle)
	obstacleTwo := new(MockObstacle)
	obstacleOne.On("IsBeyond", mock.Anything).Return(false)
	obstacleTwo.On("IsBeyond", mock.Anything).Return(false)
	return obstacleOne, obstacleTwo
}

func twoMockObstaclesOneBeyondLimit(limit size.Size) (*MockObstacle, *MockObstacle) {
	obstacleOne := new(MockObstacle)
	obstacleTwo := new(MockObstacle)
	obstacleOne.On("IsBeyond", mock.Anything).Return(true)
	obstacleTwo.On("IsBeyond", mock.Anything).Return(false)
	return obstacleOne, obstacleTwo
}

type MockObstacle struct {
	mock.Mock
	coord coordinate.Coordinate
}

func (this *MockObstacle) IsBeyond(limit size.Size) bool {
	args := this.Called(limit)
	return args.Bool(0)
}

func (this *MockObstacle) Coordinate() coordinate.Coordinate {
	return this.coord
}
