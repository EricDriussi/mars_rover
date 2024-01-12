package obstacles_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
	"mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/obstacles"
	"mars_rover/src/domain/size"
	"testing"
)

func TestDeterminesIfCoordinateIsOccupiedByAny(t *testing.T) {
	coordinate := absoluteCoordinate.Build(0, 0)
	testObstacles := setupValidObstacles(t, *coordinate)

	assert.True(t, testObstacles.Occupy(*coordinate))
	assert.False(t, testObstacles.Occupy(*absoluteCoordinate.Build(1, 1)))
}

func TestDeterminesIfAnyObstacleIsBeyondSize(t *testing.T) {
	sizeLimit, err := size.Square(2)
	assert.Nil(t, err)
	coords := absoluteCoordinate.Build(3, 3)
	testObstacles := setupValidObstacles(t, *coords)

	assert.True(t, testObstacles.IsAnyBeyond(*sizeLimit))
}

func TestDeterminesIfNoObstacleIsBeyondSize(t *testing.T) {
	sizeLimit, err := size.Square(4)
	assert.Nil(t, err)
	coords := absoluteCoordinate.Build(3, 3)
	testObstacles := setupValidObstacles(t, *coords)

	assert.False(t, testObstacles.IsAnyBeyond(*sizeLimit))
}

func TestCannotAddObstacleWithOverlappingCoordinates(t *testing.T) {
	coordinate := absoluteCoordinate.Build(0, 0)
	testObstacles := setupValidObstacles(t, *coordinate)
	coords, err := coordinates.New(
		*coordinate,
		*absoluteCoordinate.Build(0, 1),
	)
	assert.Nil(t, err)
	testObstacle, err := obstacle.CreateObstacle(*coords)
	assert.Nil(t, err)

	err = testObstacles.Add(testObstacle)

	assert.Error(t, err)
}

func TestCanAddObstacleWithNoOverlappingCoordinates(t *testing.T) {
	coordinate := absoluteCoordinate.Build(0, 0)
	testObstacles := setupValidObstacles(t, *coordinate)
	coords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 2),
	)
	assert.Nil(t, err)
	testObstacle, err := obstacle.CreateObstacle(*coords)
	assert.Nil(t, err)

	err = testObstacles.Add(testObstacle)

	assert.Nil(t, err)
}

func setupValidObstacles(t *testing.T, coord absoluteCoordinate.AbsoluteCoordinate) *obstacles.Obstacles {
	coords, err := coordinates.New(coord)
	assert.Nil(t, err)
	obstacleOne, err := obstacle.CreateObstacle(*coords)
	assert.Nil(t, err)
	testObstacles, err := obstacles.FromList(obstacleOne)
	assert.Nil(t, err)
	assert.NotNil(t, testObstacles)
	return testObstacles
}
