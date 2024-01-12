package obstacles_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
	"mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/obstacles"
	"testing"
)

func TestCannotCreateFromEmptyList(t *testing.T) {
	_, err := obstacles.FromList()

	assert.Error(t, err)
}

func TestCanCreateWithNoObstacles(t *testing.T) {
	testObstacles := obstacles.Empty()

	assert.NotNil(t, testObstacles)
	assert.Equal(t, 0, testObstacles.Amount())
}

func TestCannotCreateFromListWhenMultipleObstaclesShareCoordinates(t *testing.T) {
	coordinateOne := absoluteCoordinate.Build(0, 0)
	coordinatesOne, err := coordinates.New(*coordinateOne)
	assert.Nil(t, err)
	obstacleOne, err := obstacle.CreateObstacle(*coordinatesOne)
	assert.Nil(t, err)
	coordinateTwo := absoluteCoordinate.Build(0, 1)
	coordinatesTwo, err := coordinates.New(*coordinateOne, *coordinateTwo)
	assert.Nil(t, err)
	obstacleTwo, err := obstacle.CreateObstacle(*coordinatesTwo)

	_, err = obstacles.FromList(
		obstacleOne,
		obstacleTwo,
	)

	assert.Error(t, err)
}

func TestCanCreateFromListWhenNoObstaclesShareCoordinates(t *testing.T) {
	coordinateOne := absoluteCoordinate.Build(0, 0)
	coordinatesOne, err := coordinates.New(*coordinateOne)
	assert.Nil(t, err)
	obstacleOne, err := obstacle.CreateObstacle(*coordinatesOne)
	assert.Nil(t, err)
	coordinateTwo := absoluteCoordinate.Build(0, 1)
	coordinatesTwo, err := coordinates.New(*coordinateTwo)
	assert.Nil(t, err)
	obstacleTwo, err := obstacle.CreateObstacle(*coordinatesTwo)
	assert.Nil(t, err)

	_, err = obstacles.FromList(
		obstacleOne,
		obstacleTwo,
	)

	assert.Nil(t, err)
}
