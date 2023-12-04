package godModRover_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	rock "mars_rover/src/domain/obstacle/smallRock"
	"mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/godModRover"
	"mars_rover/src/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIgnoresCollisionMovingForward(t *testing.T) {
	planetSize, _ := size.Square(10)
	initialCoordinate := absoluteCoordinate.From(5, 5)
	obstacleCoordinate := absoluteCoordinate.From(5, 6)
	obstacleAhead := rock.In(*obstacleCoordinate)
	testPlanetWithObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{&obstacleAhead})

	testRover := godModRover.LandFacing(North{}, *initialCoordinate, testPlanetWithObstacles)

	err := testRover.MoveForward()

	assert.Equal(t, *obstacleCoordinate, testRover.Coordinate())
	assert.Nil(t, err)
}

func TestIgnoresCollisionMovingBackwards(t *testing.T) {
	planetSize, _ := size.Square(10)
	initialCoordinate := absoluteCoordinate.From(5, 5)
	obstacleCoordinate := absoluteCoordinate.From(5, 4)
	obstacleBehind := rock.In(*obstacleCoordinate)
	testPlanetWithObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{&obstacleBehind})

	testRover := godModRover.LandFacing(North{}, *initialCoordinate, testPlanetWithObstacles)

	err := testRover.MoveBackward()

	assert.Equal(t, *obstacleCoordinate, testRover.Coordinate())
	assert.Nil(t, err)
}
