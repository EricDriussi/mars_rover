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

func TestLandsOnFreeSpotFacingGivenDirection(t *testing.T) {
	testPlanet := SetupPlanetOfSizeWithObstacleIn(t, 2, *absoluteCoordinate.Build(1, 2))
	freeCoordinate := absoluteCoordinate.Build(1, 1)
	direction := North{}
	testRover := godModRover.LandFacing(uuid.New(), direction, *freeCoordinate, testPlanet)

	assert.Equal(t, *freeCoordinate, testRover.Coordinate())
	assert.Equal(t, direction, testRover.Direction())
}

func TestCanLandOnObstacle(t *testing.T) {
	coordinate := absoluteCoordinate.Build(1, 1)
	testPlanet := SetupPlanetOfSizeWithObstacleIn(t, 2, *coordinate)
	testRover := godModRover.LandFacing(uuid.New(), North{}, *coordinate, testPlanet)

	assert.NotNil(t, testRover)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}

func TestCanLandOutOfPlanet(t *testing.T) {
	testPlanet := SetupEmptyTestPlanetOfSize(t, 2)
	coordinateOutsidePlanet := absoluteCoordinate.Build(4, 3)
	testRover := godModRover.LandFacing(uuid.New(), North{}, *coordinateOutsidePlanet, testPlanet)

	assert.NotNil(t, testRover)
	assert.Equal(t, *coordinateOutsidePlanet, testRover.Coordinate())
}
