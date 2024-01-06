package planetWithObstacles_test

import (
	obs "mars_rover/src/domain/obstacle/obstacles"
	"mars_rover/src/domain/planet/planetWithObstacles"
	"mars_rover/src/domain/size"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
)

var testSize *size.Size

func init() {
	testSize, _ = size.Square(5)
}

func TestCanCreateIfNoObstacleIsOutOfBounds(t *testing.T) {
	mockObstacleOne := new(MockObstacle)
	mockObstacleTwo := new(MockObstacle)
	mockObstacleOne.On("IsBeyond", Anything).Return(false)
	mockObstacleTwo.On("IsBeyond", Anything).Return(false)
	obstaclesWithinBounds := obs.FromList(mockObstacleOne, mockObstacleTwo)

	planet, err := planetWithObstacles.Create("testColor", *testSize, *obstaclesWithinBounds)

	assert.Nil(t, err)
	assert.NotNil(t, planet)
}

func TestCannotCreateIfOneObstacleIsOutOfBounds(t *testing.T) {
	mockObstacleOne := new(MockObstacle)
	mockObstacleTwo := new(MockObstacle)
	mockObstacleOne.On("IsBeyond", Anything).Return(true)
	mockObstacleTwo.On("IsBeyond", Anything).Return(false)
	obstaclesOutsideBounds := obs.FromList(mockObstacleOne, mockObstacleTwo)

	planet, err := planetWithObstacles.Create("testColor", *testSize, *obstaclesOutsideBounds)

	assert.Error(t, err)
	assert.Nil(t, planet)
}

func TestCannotCreateIfSizeTooSmall(t *testing.T) {
	mockObstacle := new(MockObstacle)
	mockObstacle.On("IsBeyond", Anything).Return(false)
	obstacleWithinBounds := obs.FromList(mockObstacle)
	testSize, err := size.Square(1)
	assert.Nil(t, err)

	planet, err := planetWithObstacles.Create("testColor", *testSize, *obstacleWithinBounds)

	assert.Error(t, err)
	assert.Nil(t, planet)
}

func TestCannotCreateIfNoObstacles(t *testing.T) {
	planet, err := planetWithObstacles.Create("testColor", *testSize, *obs.FromList())

	assert.Error(t, err)
	assert.Nil(t, planet)
}
