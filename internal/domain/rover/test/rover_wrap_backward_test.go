package rover_test

import (
	"mars_rover/internal/domain/coordinate"
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
	landingLocation, _ := location.From(*coordinate.NewAbsolute(2, 0), &direction.North{})

	testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(*coordinate.NewAbsolute(2, 3), &direction.North{})
	assert.Equal(t, expectedLocation, testRover.Location())
}

func TestWrapsLookingEastMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(*coordinate.NewAbsolute(0, 2), &direction.East{})

	testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(*coordinate.NewAbsolute(3, 2), &direction.East{})
	assert.Equal(t, expectedLocation, testRover.Location())
}

func TestWrapsLookingSouthMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(*coordinate.NewAbsolute(2, 3), &direction.South{})

	testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(*coordinate.NewAbsolute(2, 0), &direction.South{})
	assert.Equal(t, expectedLocation, testRover.Location())
}

func TestWrapsLookingWestMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(*coordinate.NewAbsolute(3, 2), &direction.West{})

	testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(*coordinate.NewAbsolute(0, 2), &direction.West{})
	assert.Equal(t, expectedLocation, testRover.Location())
}
