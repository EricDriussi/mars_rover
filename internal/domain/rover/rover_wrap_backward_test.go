package rover_test

import (
	"mars_rover/internal/domain/coordinate"
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
	landingLocation, _ := location.From(*coordinate.New(2, 0), &direction.North{})

	testRover := rover.Land(*landingLocation, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(*coordinate.New(2, 3), &direction.North{})
	assert.True(t, expectedLocation.Equals(*testRover.Location()))
}

func TestWrapsLookingEastMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(*coordinate.New(0, 2), &direction.East{})

	testRover := rover.Land(*landingLocation, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(*coordinate.New(3, 2), &direction.East{})
	assert.True(t, expectedLocation.Equals(*testRover.Location()))
}

func TestWrapsLookingSouthMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(*coordinate.New(2, 3), &direction.South{})

	testRover := rover.Land(*landingLocation, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(*coordinate.New(2, 0), &direction.South{})
	assert.True(t, expectedLocation.Equals(*testRover.Location()))
}

func TestWrapsLookingWestMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(*coordinate.New(3, 2), &direction.West{})

	testRover := rover.Land(*landingLocation, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(*coordinate.New(0, 2), &direction.West{})
	assert.True(t, expectedLocation.Equals(*testRover.Location()))
}
