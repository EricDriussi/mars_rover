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

func TestWrapsLookingNorthMovingForward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingPosition, _ := location.From(2, 3)

	testRover := rover.Land(*landingPosition, &direction.North{}, *testPlanetWithoutObstacles)

	testRover.MoveForward()

	expectedPosition, _ := location.From(2, 0)
	assert.True(t, expectedPosition.Equals(testRover.Location()))
}

func TestWrapsLookingEastMovingForward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingPosition, _ := location.From(3, 2)

	testRover := rover.Land(*landingPosition, &direction.East{}, *testPlanetWithoutObstacles)

	testRover.MoveForward()

	expectedPosition, _ := location.From(0, 2)
	assert.True(t, expectedPosition.Equals(testRover.Location()))
}

func TestWrapsLookingSouthMovingForward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingPosition, _ := location.From(2, 0)

	testRover := rover.Land(*landingPosition, &direction.South{}, *testPlanetWithoutObstacles)

	testRover.MoveForward()

	expectedPosition, _ := location.From(2, 3)
	assert.True(t, expectedPosition.Equals(testRover.Location()))
}

func TestWrapsLookingWestMovingForward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingPosition, _ := location.From(0, 2)

	testRover := rover.Land(*landingPosition, &direction.West{}, *testPlanetWithoutObstacles)

	testRover.MoveForward()

	expectedPosition, _ := location.From(3, 2)
	assert.True(t, expectedPosition.Equals(testRover.Location()))
}
