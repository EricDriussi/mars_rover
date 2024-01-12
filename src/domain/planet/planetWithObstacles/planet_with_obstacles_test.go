package planetWithObstacles_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
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
	mockCoordinatesOne, err := coordinates.New(*absoluteCoordinate.Build(0, 0))
	assert.Nil(t, err)
	mockObstacleOne.On("Coordinates").Return(*mockCoordinatesOne)
	mockCoordinatesTwo, err := coordinates.New(*absoluteCoordinate.Build(0, 1))
	assert.Nil(t, err)
	mockObstacleTwo.On("Coordinates").Return(*mockCoordinatesTwo)
	mockObstacleOne.On("IsBeyond", Anything).Return(false)
	mockObstacleTwo.On("IsBeyond", Anything).Return(false)
	obstaclesWithinBounds, err := obs.FromList(mockObstacleOne, mockObstacleTwo)
	assert.Nil(t, err)

	planet, err := planetWithObstacles.Create("testColor", *testSize, *obstaclesWithinBounds)

	assert.Nil(t, err)
	assert.NotNil(t, planet)
}

func TestCannotCreateIfOneObstacleIsOutOfBounds(t *testing.T) {
	mockObstacleOne := new(MockObstacle)
	mockObstacleTwo := new(MockObstacle)
	mockCoordinatesOne, err := coordinates.New(*absoluteCoordinate.Build(0, 0))
	assert.Nil(t, err)
	mockObstacleOne.On("Coordinates").Return(*mockCoordinatesOne)
	mockCoordinatesTwo, err := coordinates.New(*absoluteCoordinate.Build(0, 1))
	assert.Nil(t, err)
	mockObstacleTwo.On("Coordinates").Return(*mockCoordinatesTwo)
	mockObstacleOne.On("IsBeyond", Anything).Return(true)
	mockObstacleTwo.On("IsBeyond", Anything).Return(false)
	obstaclesOutsideBounds, err := obs.FromList(mockObstacleOne, mockObstacleTwo)
	assert.Nil(t, err)

	_, err = planetWithObstacles.Create("testColor", *testSize, *obstaclesOutsideBounds)

	assert.Error(t, err)
}

func TestCannotCreateIfSizeTooSmall(t *testing.T) {
	mockObstacle := new(MockObstacle)
	mockCoordinates, err := coordinates.New(*absoluteCoordinate.Build(0, 0))
	mockObstacle.On("Coordinates").Return(*mockCoordinates)
	mockObstacle.On("IsBeyond", Anything).Return(false)
	obstacleWithinBounds, err := obs.FromList(mockObstacle)
	assert.Nil(t, err)
	testSize, err := size.Square(1)
	assert.Nil(t, err)

	_, err = planetWithObstacles.Create("testColor", *testSize, *obstacleWithinBounds)

	assert.Error(t, err)
}

func TestCannotCreateIfNoObstacles(t *testing.T) {
	_, err := planetWithObstacles.Create("testColor", *testSize, *obs.Empty())

	assert.Error(t, err)
}
