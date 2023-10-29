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

func TestAvoidsCollisionMovingForward(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	landingLocation, _ := location.From(5, 5)
	obstacleLocation, _ := location.From(5, 6)
	obstacleInfront := obstacle.In(obstacleLocation)
	testPlanetWithObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{*obstacleInfront})

	testRover := rover.Land(*landingLocation, &direction.North{}, *testPlanetWithObstacles)

	testRover.MoveForward()

	expectedLocation, _ := location.From(5, 5)
	assert.True(t, expectedLocation.Equals(testRover.Location()))
}

func TestAvoidsCollisionWrappingForward(t *testing.T) {
	planetSize, _ := size.From(5, 5)
	landingLocation, _ := location.From(3, 5)
	obstacleLocation, _ := location.From(3, 0)
	obstacleInfront := obstacle.In(obstacleLocation)
	testPlanetWithObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{*obstacleInfront})

	testRover := rover.Land(*landingLocation, &direction.North{}, *testPlanetWithObstacles)

	testRover.MoveForward()

	expectedLocation, _ := location.From(3, 5)
	assert.True(t, expectedLocation.Equals(testRover.Location()))
}

func TestAvoidsCollisionMovingBackwards(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	landingLocation, _ := location.From(5, 5)
	obstacleLocation, _ := location.From(5, 4)
	obstacleBehind := obstacle.In(obstacleLocation)
	testPlanetWithObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{*obstacleBehind})

	testRover := rover.Land(*landingLocation, &direction.North{}, *testPlanetWithObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(5, 5)
	assert.True(t, expectedLocation.Equals(testRover.Location()))
}

func TestAvoidsCollisionWrappingBackwards(t *testing.T) {
	planetSize, _ := size.From(5, 5)
	landingLocation, _ := location.From(3, 0)
	obstacleLocation, _ := location.From(3, 5)
	obstacleBehind := obstacle.In(obstacleLocation)
	testPlanetWithObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{*obstacleBehind})

	testRover := rover.Land(*landingLocation, &direction.North{}, *testPlanetWithObstacles)

	testRover.MoveBackward()

	expectedLocation, _ := location.From(3, 0)
	assert.True(t, expectedLocation.Equals(testRover.Location()))
}
