package infra_test

import (
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
