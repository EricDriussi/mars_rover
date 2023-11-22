package infra_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/bigRock"
	"mars_rover/internal/domain/obstacle/smallRock"
	. "mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/planet/rockyPlanet"
	. "mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/rover/godModRover"
	"mars_rover/internal/domain/rover/wrappingCollidingRover"
	s "mars_rover/internal/domain/size"
	"testing"
)

func setupWrappingRoverOnRockyPlanet() (Rover, Planet) {
	rovCoord := absoluteCoordinate.From(0, 0)
	testPlanet := setupRockyPlanet()
	testRover, _ := wrappingCollidingRover.Land(*rovCoord, testPlanet)
	return testRover, testPlanet
}

func setupGodModRoverOnRockyPlanet() (Rover, Planet) {
	rovCoord := absoluteCoordinate.From(1, 1)
	testPlanet := setupRockyPlanet()
	testRover := godModRover.Land(*rovCoord, testPlanet)
	return testRover, testPlanet
}

func setupRockyPlanet() Planet {
	size, _ := s.Square(10)
	smallCoord := absoluteCoordinate.From(1, 1)
	testSmallRock := smallRock.In(*smallCoord)
	bigCoord1 := absoluteCoordinate.From(2, 2)
	bigCoord2 := absoluteCoordinate.From(2, 3)
	testBigRock := bigRock.In([]absoluteCoordinate.AbsoluteCoordinate{*bigCoord1, *bigCoord2})
	testPlanet, _ := rockyPlanet.Create("testColor", *size, []Obstacle{&testSmallRock, &testBigRock})
	return testPlanet
}

func assertPlanetsAreEqual(t *testing.T, testPlanet Planet, foundPlanet Planet) {
	assert.Equal(t, testPlanet.Color(), foundPlanet.Color())
	assert.Equal(t, testPlanet.Obstacles(), foundPlanet.Obstacles())
	assert.Equal(t, testPlanet.Size(), foundPlanet.Size())
}

func assertRoversAreEqual(t *testing.T, foundRover Rover, testRover Rover) {
	assert.Equal(t, testRover.Coordinate(), foundRover.Coordinate())
	assert.Equal(t, testRover.Direction().CardinalPoint(), foundRover.Direction().CardinalPoint())
	assert.Equal(t, testRover.Map(), foundRover.Map())
}
