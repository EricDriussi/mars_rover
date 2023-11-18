package wrappingCollidingRover_test

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/direction"
	. "mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/smallRock"
	"mars_rover/internal/domain/planet/rockyPlanet"
	"mars_rover/internal/domain/rover/wrappingCollidingRover"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLandsOnFreeSpot(t *testing.T) {
	planetSize, _ := size.Square(2)
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{smallRock.In(*absoluteCoordinate.From(1, 2))})
	coordinate := absoluteCoordinate.From(1, 1)

	testRover, err := wrappingCollidingRover.Land(*coordinate, testPlanet)

	assert.Nil(t, err)
	assert.Equal(t, *coordinate, testRover.Coordinate())
}

func TestLandsOnFreeSpotFacingGivenDirection(t *testing.T) {
	planetSize, _ := size.Square(2)
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{smallRock.In(*absoluteCoordinate.From(1, 2))})
	coordinate := absoluteCoordinate.From(1, 1)
	direction := North{}

	testRover, err := wrappingCollidingRover.LandFacing(direction, *coordinate, testPlanet)

	assert.Nil(t, err)
	assert.Equal(t, *coordinate, testRover.Coordinate())
	assert.Equal(t, direction, testRover.Direction())
}

func TestCannotLandOnObstacle(t *testing.T) {
	planetSize, _ := size.Square(2)
	coordinate := absoluteCoordinate.From(1, 1)
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{smallRock.In(*coordinate)})

	testRover, err := wrappingCollidingRover.Land(*coordinate, testPlanet)

	assert.Error(t, err)
	assert.Nil(t, testRover)
}
