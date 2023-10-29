package rover_test

import (
	"mars_rover/internal/domain/direction"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapsLookingNorthMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(2, 0)

	testRover := rover.Land(*landingLocation, &direction.North{}, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(2, 3)
	assert.True(t, expectedLocation.Equals(testRover.Location()))
}

func TestWrapsLookingEastMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(0, 2)

	testRover := rover.Land(*landingLocation, &direction.East{}, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(3, 2)
	assert.True(t, expectedLocation.Equals(testRover.Location()))
}

func TestWrapsLookingSouthMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(2, 3)

	testRover := rover.Land(*landingLocation, &direction.South{}, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(2, 0)
	assert.True(t, expectedLocation.Equals(testRover.Location()))
}

func TestWrapsLookingWestMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(3, 2)

	testRover := rover.Land(*landingLocation, &direction.West{}, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(0, 2)
	assert.True(t, expectedLocation.Equals(testRover.Location()))
}
