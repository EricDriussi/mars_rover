package rover_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	rock "mars_rover/internal/domain/obstacle/small_rock"
	rockyPlanet "mars_rover/internal/domain/planet/rocky_planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvoidsCollisionMovingForward(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	landingLocation, _ := location.From(*coordinate.New(5, 5), &direction.North{})
	obstacleInfront := rock.In(*coordinate.New(5, 6))
	testPlanetWithObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{obstacleInfront})

	testRover := rover.Land(*landingLocation, testPlanetWithObstacles)

	testRover.MoveForward()

	assert.True(t, landingLocation.Equals(*testRover.Location()))
}

func TestAvoidsCollisionWrappingForward(t *testing.T) {
	planetSize, _ := size.From(5, 5)
	landingLocation, _ := location.From(*coordinate.New(3, 5), &direction.North{})
	obstacleLocation := *coordinate.New(3, 0)
	obstacleInfront := rock.In(obstacleLocation)
	testPlanetWithObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{obstacleInfront})

	testRover := rover.Land(*landingLocation, testPlanetWithObstacles)

	testRover.MoveForward()

	assert.True(t, landingLocation.Equals(*testRover.Location()))
}

func TestAvoidsCollisionMovingBackwards(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	landingLocation, _ := location.From(*coordinate.New(5, 5), &direction.North{})
	obstacleLocation := *coordinate.New(5, 4)
	obstacleBehind := rock.In(obstacleLocation)
	testPlanetWithObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{obstacleBehind})

	testRover := rover.Land(*landingLocation, testPlanetWithObstacles)

	testRover.MoveBackward()

	assert.True(t, landingLocation.Equals(*testRover.Location()))
}

func TestAvoidsCollisionWrappingBackwards(t *testing.T) {
	planetSize, _ := size.From(5, 5)
	landingLocation, _ := location.From(*coordinate.New(3, 0), &direction.North{})
	obstacleLocation := *coordinate.New(3, 5)
	obstacleBehind := rock.In(obstacleLocation)
	testPlanetWithObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{obstacleBehind})

	testRover := rover.Land(*landingLocation, testPlanetWithObstacles)

	testRover.MoveBackward()

	assert.True(t, landingLocation.Equals(*testRover.Location()))
}
