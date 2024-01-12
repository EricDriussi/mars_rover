package infra_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
	"mars_rover/src/domain/obstacle"
	obs "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/planet"
	"mars_rover/src/domain/planet/planetWithObstacles"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/godModRover"
	"mars_rover/src/domain/rover/id"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	s "mars_rover/src/domain/size"
	"testing"
)

// TODO.LM: These helpers have a lot of hard coded values.
// It is ok in this case, since as long as the returned objects are valid, the values don't matter.
// This might indicate that mocking the objects would be a better approach.
// However, since these are used to test persistence, I think avoiding mocks gives more assurance.

func setupWrappingRoverOnRockyPlanet(t *testing.T) (Rover, Planet) {
	rovCoord := absoluteCoordinate.Build(0, 0)
	testPlanet := setupPlanet(t)
	aDirection := North{}
	testRover, err := wrappingCollidingRover.LandFacing(id.New(), aDirection, *rovCoord, testPlanet)
	assert.Nil(t, err)
	return testRover, testPlanet
}

func setupGodModRoverOnRockyPlanet(t *testing.T) (Rover, Planet) {
	rovCoord := absoluteCoordinate.Build(1, 1)
	testPlanet := setupPlanet(t)
	aDirection := North{}
	testRover := godModRover.LandFacing(id.New(), aDirection, *rovCoord, testPlanet)
	return testRover, testPlanet
}

func setupPlanet(t *testing.T) Planet {
	size, err := s.Square(10)
	assert.Nil(t, err)
	smallCoords, err := coordinates.New(*absoluteCoordinate.Build(1, 1))
	assert.Nil(t, err)
	testSmall, err := obstacle.CreateObstacle(*smallCoords)
	assert.Nil(t, err)
	bigCoords, err := coordinates.New(*absoluteCoordinate.Build(2, 2), *absoluteCoordinate.Build(2, 3))
	assert.Nil(t, err)
	testBig, err := obstacle.CreateObstacle(*bigCoords)
	assert.Nil(t, err)
	obstacles, err := obs.FromList(testSmall, testBig)
	assert.Nil(t, err)
	testPlanet, err := planetWithObstacles.Create("testColor", *size, *obstacles)
	assert.Nil(t, err)
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
	// TODO.LM: Normally, this comparison would suggest the need for an Equals method in the Direction interface.
	// However, that would mean implementing a function in prod code that would only be used in tests.
	assert.Equal(t, testRover.Direction().CardinalPoint(), foundRover.Direction().CardinalPoint())
	assert.Equal(t, testRover.Map(), foundRover.Map())
}
