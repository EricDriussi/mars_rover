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
	testCoordinate := absoluteCoordinate.Build(0, 0)
	testObstacles := setupValidObstacles(t, *testCoordinate)

	assert.True(t, testObstacles.Occupy(*testCoordinate))
	assert.False(t, testObstacles.Occupy(*absoluteCoordinate.Build(1, 1)))
}

func TestDeterminesIfAnyObstacleIsBeyondSize(t *testing.T) {
	testSize, err := size.Square(2)
	assert.Nil(t, err)
	testCoordinates := absoluteCoordinate.Build(3, 3)
	testObstacles := setupValidObstacles(t, *testCoordinates)

	assert.True(t, testObstacles.IsAnyBeyond(*testSize))
}

func TestDeterminesIfNoObstacleIsBeyondSize(t *testing.T) {
	testSize, err := size.Square(4)
	assert.Nil(t, err)
	testCoordinates := absoluteCoordinate.Build(3, 3)
	testObstacles := setupValidObstacles(t, *testCoordinates)

	assert.False(t, testObstacles.IsAnyBeyond(*testSize))
}

func TestCannotAddObstacleWithOverlappingCoordinates(t *testing.T) {
	testCoordinate := absoluteCoordinate.Build(0, 0)
	testObstacles := setupValidObstacles(t, *testCoordinate)
	testCoordinates, err := coordinates.New(
		*testCoordinate,
		*absoluteCoordinate.Build(0, 1),
	)
	assert.Nil(t, err)
	testObstacle, err := obstacle.CreateObstacle(*testCoordinates)
	assert.Nil(t, err)

	err = testObstacles.Add(testObstacle)

	assert.Error(t, err)
}

func TestCanAddObstacleWithNoOverlappingCoordinates(t *testing.T) {
	testCoordinate := absoluteCoordinate.Build(0, 0)
	testObstacles := setupValidObstacles(t, *testCoordinate)
	testCoordinates, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 2),
	)
	assert.Nil(t, err)
	testObstacle, err := obstacle.CreateObstacle(*testCoordinates)
	assert.Nil(t, err)

	err = testObstacles.Add(testObstacle)

	assert.Nil(t, err)
}

func setupValidObstacles(t *testing.T, coord absoluteCoordinate.AbsoluteCoordinate) *obstacles.Obstacles {
	testCoordinates, err := coordinates.New(coord)
	assert.Nil(t, err)
	obstacleOne, err := obstacle.CreateObstacle(*testCoordinates)
	assert.Nil(t, err)
	testObstacles, err := obstacles.FromList(obstacleOne)
	assert.Nil(t, err)
	assert.NotNil(t, testObstacles)
	return testObstacles
}
