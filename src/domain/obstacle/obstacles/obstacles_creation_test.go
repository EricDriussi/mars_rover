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
	testObstacles, err := obstacles.FromList()

	assert.Error(t, err)
	assert.Nil(t, testObstacles)
}

func TestCanCreateWithNoObstacles(t *testing.T) {
	testObstacles := obstacles.Empty()

	assert.NotNil(t, testObstacles)
	assert.Equal(t, 0, testObstacles.Amount())
}

func TestCannotCreateFromListWhenMultipleObstaclesShareCoordinates(t *testing.T) {
	coordOne := absoluteCoordinate.Build(0, 0)
	coordsOne, err := coordinates.New(*coordOne)
	assert.Nil(t, err)
	obstacleOne, err := obstacle.CreateObstacle(*coordsOne)
	assert.Nil(t, err)
	coordTwo := absoluteCoordinate.Build(0, 1)
	coordsTwo, err := coordinates.New(*coordOne, *coordTwo)
	assert.Nil(t, err)
	obstacleTwo, err := obstacle.CreateObstacle(*coordsTwo)

	testObstacles, err := obstacles.FromList(
		obstacleOne,
		obstacleTwo,
	)

	assert.Error(t, err)
	assert.Nil(t, testObstacles)
}

func TestCanCreateFromListWhenNoObstaclesShareCoordinates(t *testing.T) {
	coordOne := absoluteCoordinate.Build(0, 0)
	coordsOne, err := coordinates.New(*coordOne)
	assert.Nil(t, err)
	obstacleOne, err := obstacle.CreateObstacle(*coordsOne)
	assert.Nil(t, err)
	coordTwo := absoluteCoordinate.Build(0, 1)
	coordsTwo, err := coordinates.New(*coordTwo)
	assert.Nil(t, err)
	obstacleTwo, err := obstacle.CreateObstacle(*coordsTwo)
	assert.Nil(t, err)

	testObstacles, err := obstacles.FromList(
		obstacleOne,
		obstacleTwo,
	)

	assert.Nil(t, err)
	assert.NotNil(t, testObstacles)
}
