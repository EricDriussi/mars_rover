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

func TestLandsOnFreeSpotFacingGivenDirection(t *testing.T) {
	testPlanet := SetupPlanetOfSizeWithObstacleIn(t, 2, *absoluteCoordinate.Build(1, 2))
	direction := North{}
	coordinate := absoluteCoordinate.Build(1, 1)
	testRover, err := wrappingCollidingRover.LandFacing(uuid.New(), direction, *coordinate, testPlanet)

	assert.Nil(t, err)
	assert.Equal(t, direction, testRover.Direction())
	assert.Equal(t, *coordinate, testRover.Coordinate())
}

func TestCannotLandOnObstacle(t *testing.T) {
	coordinate := absoluteCoordinate.Build(1, 1)
	testPlanet := SetupPlanetOfSizeWithObstacleIn(t, 2, *coordinate)

	direction := North{}
	testRover, err := wrappingCollidingRover.LandFacing(uuid.New(), direction, *coordinate, testPlanet)

	assert.Error(t, err)
	assert.Nil(t, testRover)
}

func TestCannotLandOutOfPlanet(t *testing.T) {
	testPlanet := SetupEmptyTestPlanetOfSize(t, 2)

	testRover, err := wrappingCollidingRover.LandFacing(uuid.New(), North{}, *absoluteCoordinate.Build(4, 3), testPlanet)

	assert.Error(t, err)
	assert.Nil(t, testRover)
}
