package godModRover_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/godModRover"
	"mars_rover/src/domain/rover/uuid"
	. "mars_rover/src/test_helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIgnoresCollisionMovingForward(t *testing.T) {
	obstacleCoordinate := absoluteCoordinate.Build(5, 6)
	testPlanet := SetupPlanetOfSizeWithObstacleIn(t, 10, *obstacleCoordinate)
	testRover := godModRover.LandFacing(uuid.New(), North{}, *absoluteCoordinate.Build(5, 5), testPlanet)

	err := testRover.MoveForward()

	assert.Equal(t, *obstacleCoordinate, testRover.Coordinate())
	assert.Nil(t, err)
}

func TestIgnoresCollisionMovingBackwards(t *testing.T) {
	obstacleCoordinate := absoluteCoordinate.Build(5, 4)
	testPlanet := SetupPlanetOfSizeWithObstacleIn(t, 10, *obstacleCoordinate)
	testRover := godModRover.LandFacing(uuid.New(), North{}, *absoluteCoordinate.Build(5, 5), testPlanet)

	err := testRover.MoveBackward()

	assert.Equal(t, *obstacleCoordinate, testRover.Coordinate())
	assert.Nil(t, err)
}
