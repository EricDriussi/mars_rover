package infra_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/bigRock"
	"mars_rover/src/domain/obstacle/smallRock"
	. "mars_rover/src/domain/planet"
	"mars_rover/src/domain/planet/planetWithObstacles"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/godModRover"
	"mars_rover/src/domain/rover/uuid"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	s "mars_rover/src/domain/size"
	"testing"
)

// TODO: Do something with these hardcoded values and error handling

func setupWrappingRoverOnRockyPlanet() (Rover, Planet) {
	rovCoord := absoluteCoordinate.Build(0, 0)
	testPlanet := setupRockyPlanet()
	aDirection := North{}
	testRover, _ := wrappingCollidingRover.LandFacing(uuid.New(), aDirection, *rovCoord, testPlanet)
	return testRover, testPlanet
}

func setupGodModRoverOnRockyPlanet() (Rover, Planet) {
	rovCoord := absoluteCoordinate.Build(1, 1)
	testPlanet := setupRockyPlanet()
	aDirection := North{}
	testRover := godModRover.LandFacing(uuid.New(), aDirection, *rovCoord, testPlanet)
	return testRover, testPlanet
}

func setupRockyPlanet() Planet {
	size, _ := s.Square(10)
	smallCoord := absoluteCoordinate.Build(1, 1)
	testSmallRock := smallRock.In(*smallCoord)
	bigCoord1 := absoluteCoordinate.Build(2, 2)
	bigCoord2 := absoluteCoordinate.Build(2, 3)
	testBigRock, _ := bigRock.In(*bigCoord1, *bigCoord2)
	testPlanet, _ := planetWithObstacles.Create("testColor", *size, []Obstacle{&testSmallRock, testBigRock})
	return testPlanet
}

func assertPlanetsAreEqual(t *testing.T, testPlanet Planet, foundPlanet Planet) {
	assert.Equal(t, testPlanet.Color(), foundPlanet.Color())
	assert.Equal(t, testPlanet.Obstacles(), foundPlanet.Obstacles())
	assert.Equal(t, testPlanet.Size(), foundPlanet.Size())
}

func assertRoversAreEqual(t *testing.T, foundRover Rover, testRover Rover) {
	assert.Equal(t, testRover.Id(), foundRover.Id())
	assert.Equal(t, testRover.Coordinate(), foundRover.Coordinate())
	assert.Equal(t, testRover.Direction().CardinalPoint(), foundRover.Direction().CardinalPoint())
	assert.Equal(t, testRover.Map(), foundRover.Map())
}
