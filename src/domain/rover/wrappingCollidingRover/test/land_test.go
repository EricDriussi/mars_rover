package wrappingCollidingRover_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/smallRock"
	"mars_rover/src/domain/planet/planetWithObstacles"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/uuid"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	"mars_rover/src/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLandsOnFreeSpotFacingGivenDirection(t *testing.T) {
	planetSize, _ := size.Square(2)
	rock := smallRock.In(*absoluteCoordinate.Build(1, 2))
	testPlanet, _ := planetWithObstacles.Create("testColor", *planetSize, []Obstacle{&rock})
	coordinate := absoluteCoordinate.Build(1, 1)
	direction := North{}

	testRover, err := wrappingCollidingRover.LandFacing(uuid.New(), direction, *coordinate, testPlanet)

	assert.Nil(t, err)
	assert.Equal(t, *coordinate, testRover.Coordinate())
	assert.Equal(t, direction, testRover.Direction())
}

func TestCannotLandOnObstacle(t *testing.T) {
	planetSize, _ := size.Square(2)
	coordinate := absoluteCoordinate.Build(1, 1)
	rock := smallRock.In(*coordinate)
	testPlanet, _ := planetWithObstacles.Create("testColor", *planetSize, []Obstacle{&rock})
	aDirection := North{}

	testRover, err := wrappingCollidingRover.LandFacing(uuid.New(), aDirection, *coordinate, testPlanet)

	assert.Error(t, err)
	assert.Nil(t, testRover)
}

func TestCannotLandOutOfPlanet(t *testing.T) {
	planetSize, _ := size.Square(2)
	coordinate := absoluteCoordinate.Build(4, 3)
	rock := smallRock.In(*absoluteCoordinate.Build(1, 1))
	testPlanet, _ := planetWithObstacles.Create("testColor", *planetSize, []Obstacle{&rock})
	aDirection := North{}

	testRover, err := wrappingCollidingRover.LandFacing(uuid.New(), aDirection, *coordinate, testPlanet)

	assert.Error(t, err)
	assert.Nil(t, testRover)
}
