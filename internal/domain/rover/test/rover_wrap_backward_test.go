package rover_test

import (
	coordinate2d "mars_rover/internal/domain/coordinate/coordinate2D"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	rockyPlanet "mars_rover/internal/domain/planet/rocky_planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapsLookingNorthMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(coordinate2d.New(2, 0), &direction.North{})

	testRover := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(coordinate2d.New(2, 3), &direction.North{})
	assert.True(t, expectedLocation.Equals(*testRover.Location()))
}

func TestWrapsLookingEastMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(coordinate2d.New(0, 2), &direction.East{})

	testRover := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(coordinate2d.New(3, 2), &direction.East{})
	assert.True(t, expectedLocation.Equals(*testRover.Location()))
}

func TestWrapsLookingSouthMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(coordinate2d.New(2, 3), &direction.South{})

	testRover := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(coordinate2d.New(2, 0), &direction.South{})
	assert.True(t, expectedLocation.Equals(*testRover.Location()))
}

func TestWrapsLookingWestMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(coordinate2d.New(3, 2), &direction.West{})

	testRover := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(coordinate2d.New(0, 2), &direction.West{})
	assert.True(t, expectedLocation.Equals(*testRover.Location()))
}
