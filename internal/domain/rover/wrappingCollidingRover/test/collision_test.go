package wrappingCollidingRover_test

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle"
	rock "mars_rover/internal/domain/obstacle/smallRock"
	"mars_rover/internal/domain/planet/rockyPlanet"
	. "mars_rover/internal/domain/rover/direction"
	"mars_rover/internal/domain/rover/wrappingCollidingRover"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvoidsCollisionMovingForward(t *testing.T) {
	planetSize, _ := size.Square(10)
	coordinate := absoluteCoordinate.From(5, 5)
	obstacleAhead := rock.In(*absoluteCoordinate.From(5, 6))
	testPlanetWithObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{obstacleAhead})

	testRover, _ := wrappingCollidingRover.LandFacing(North{}, *coordinate, testPlanetWithObstacles)

	err := testRover.MoveForward()

	assert.Error(t, err)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}

func TestAvoidsCollisionWrappingForward(t *testing.T) {
	planetSize, _ := size.Square(5)
	coordinate := absoluteCoordinate.From(3, 5)
	obstacleAhead := rock.In(*absoluteCoordinate.From(3, 0))
	testPlanetWithObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{obstacleAhead})

	testRover, _ := wrappingCollidingRover.LandFacing(North{}, *coordinate, testPlanetWithObstacles)

	err := testRover.MoveForward()

	assert.Error(t, err)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}

func TestAvoidsCollisionMovingBackwards(t *testing.T) {
	planetSize, _ := size.Square(10)
	coordinate := absoluteCoordinate.From(5, 5)
	obstacleBehind := rock.In(*absoluteCoordinate.From(5, 4))
	testPlanetWithObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{obstacleBehind})

	testRover, _ := wrappingCollidingRover.LandFacing(North{}, *coordinate, testPlanetWithObstacles)

	err := testRover.MoveBackward()

	assert.Error(t, err)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}

func TestAvoidsCollisionWrappingBackwards(t *testing.T) {
	planetSize, _ := size.Square(5)
	coordinate := absoluteCoordinate.From(3, 0)
	obstacleBehind := rock.In(*absoluteCoordinate.From(3, 5))
	testPlanetWithObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{obstacleBehind})

	testRover, _ := wrappingCollidingRover.LandFacing(North{}, *coordinate, testPlanetWithObstacles)

	err := testRover.MoveBackward()

	assert.Error(t, err)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}
