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

func TestWrapsLookingNorthMovingForward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(*coordinate.NewAbsolute(2, 3), &direction.North{})

	testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	err := testRover.MoveForward()

	expectedLocation, _ := location.From(*coordinate.NewAbsolute(2, 0), &direction.North{})
	assert.Nil(t, err)
	assert.Equal(t, expectedLocation, testRover.Location())
}

func TestWrapsLookingEastMovingForward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(*coordinate.NewAbsolute(3, 2), &direction.East{})

	testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	err := testRover.MoveForward()

	expectedLocation, _ := location.From(*coordinate.NewAbsolute(0, 2), &direction.East{})
	assert.Nil(t, err)
	assert.Equal(t, expectedLocation, testRover.Location())
}

func TestWrapsLookingSouthMovingForward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(*coordinate.NewAbsolute(2, 0), &direction.South{})

	testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	err := testRover.MoveForward()

	expectedLocation, _ := location.From(*coordinate.NewAbsolute(2, 3), &direction.South{})
	assert.Nil(t, err)
	assert.Equal(t, expectedLocation, testRover.Location())
}

func TestWrapsLookingWestMovingForward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	landingLocation, _ := location.From(*coordinate.NewAbsolute(0, 2), &direction.West{})

	testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

	err := testRover.MoveForward()

	expectedLocation, _ := location.From(*coordinate.NewAbsolute(3, 2), &direction.West{})
	assert.Nil(t, err)
	assert.Equal(t, expectedLocation, testRover.Location())
}
