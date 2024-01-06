package wrappingCollidingRover_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/uuid"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	. "mars_rover/src/test_helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvoidsCollisionMovingForward(t *testing.T) {
	testPlanet := SetupPlanetOfSizeWithObstacleIn(t, 10, *absoluteCoordinate.Build(5, 6))
	coordinate := absoluteCoordinate.Build(5, 5)
	testRover, err := wrappingCollidingRover.LandFacing(uuid.New(), North{}, *coordinate, testPlanet)
	assert.Nil(t, err)

	err = testRover.MoveForward()

	assert.Error(t, err)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}

func TestAvoidsCollisionWrappingForward(t *testing.T) {
	testPlanet := SetupPlanetOfSizeWithObstacleIn(t, 6, *absoluteCoordinate.Build(3, 0))
	coordinate := absoluteCoordinate.Build(3, 5)
	testRover, err := wrappingCollidingRover.LandFacing(uuid.New(), North{}, *coordinate, testPlanet)
	assert.Nil(t, err)

	err = testRover.MoveForward()

	assert.Error(t, err)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}

func TestAvoidsCollisionMovingBackwards(t *testing.T) {
	testPlanet := SetupPlanetOfSizeWithObstacleIn(t, 10, *absoluteCoordinate.Build(5, 4))
	coordinate := absoluteCoordinate.Build(5, 5)
	testRover, err := wrappingCollidingRover.LandFacing(uuid.New(), North{}, *coordinate, testPlanet)
	assert.Nil(t, err)

	err = testRover.MoveBackward()

	assert.Error(t, err)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}

func TestAvoidsCollisionWrappingBackwards(t *testing.T) {
	testPlanet := SetupPlanetOfSizeWithObstacleIn(t, 6, *absoluteCoordinate.Build(3, 5))
	coordinate := absoluteCoordinate.Build(3, 0)
	testRover, err := wrappingCollidingRover.LandFacing(uuid.New(), North{}, *coordinate, testPlanet)
	assert.Nil(t, err)

	err = testRover.MoveBackward()

	assert.Error(t, err)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}
