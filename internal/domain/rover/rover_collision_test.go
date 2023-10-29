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

// TODO: what if there is an obstacle after wrapping?

func TestAvoidsCollisionMovingForward(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	landingLocation, _ := location.From(5, 5)
	obstacleLocation, _ := location.From(5, 6)
	obstacleInfront := obstacle.In(obstacleLocation)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{*obstacleInfront})

	testRover := rover.Land(*landingLocation, &direction.North{}, *testPlanetWithoutObstacles)

	testRover.MoveForward()

	expectedLocation, _ := location.From(5, 5)
	assert.True(t, expectedLocation.Equals(testRover.Location()))
}

func TestAvoidsCollisionMovingBackwards(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	landingLocation, _ := location.From(5, 5)
	obstacleLocation, _ := location.From(5, 4)
	obstacleInfront := obstacle.In(obstacleLocation)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{*obstacleInfront})

	testRover := rover.Land(*landingLocation, &direction.North{}, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(5, 5)
	assert.True(t, expectedLocation.Equals(testRover.Location()))
}
