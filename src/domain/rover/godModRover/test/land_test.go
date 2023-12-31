package godModRover_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/smallRock"
	"mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/godModRover"
	"mars_rover/src/domain/rover/uuid"
	"mars_rover/src/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLandsOnFreeSpotFacingGivenDirection(t *testing.T) {
	planetSize, _ := size.Square(2)
	rock := smallRock.In(*absoluteCoordinate.From(1, 2))
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{&rock})
	coordinate := absoluteCoordinate.From(1, 1)
	direction := North{}

	testRover := godModRover.LandFacing(uuid.New(), direction, *coordinate, testPlanet)

	assert.Equal(t, *coordinate, testRover.Coordinate())
	assert.Equal(t, direction, testRover.Direction())
}

func TestCanLandOnObstacle(t *testing.T) {
	planetSize, _ := size.Square(2)
	coordinate := absoluteCoordinate.From(1, 1)
	rock := smallRock.In(*coordinate)
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{&rock})
	direction := North{}

	testRover := godModRover.LandFacing(uuid.New(), direction, *coordinate, testPlanet)

	assert.NotNil(t, testRover)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}

func TestCannotLandOutOfPlanet(t *testing.T) {
	planetSize, _ := size.Square(2)
	coordinate := absoluteCoordinate.From(4, 3)
	rock := smallRock.In(*absoluteCoordinate.From(1, 1))
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{&rock})
	direction := North{}

	testRover := godModRover.LandFacing(uuid.New(), direction, *coordinate, testPlanet)

	assert.NotNil(t, testRover)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}
