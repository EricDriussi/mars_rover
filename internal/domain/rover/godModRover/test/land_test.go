package godModRover_test

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/direction"
	. "mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/smallRock"
	"mars_rover/internal/domain/planet/rockyPlanet"
	"mars_rover/internal/domain/rover/godModRover"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLandsOnFreeLocation(t *testing.T) {
	planetSize, _ := size.Square(2)
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{smallRock.In(*absoluteCoordinate.From(1, 2))})
	coordinate := absoluteCoordinate.From(1, 1)

	testRover := godModRover.Land(*coordinate, testPlanet)

	assert.Equal(t, *coordinate, testRover.Position())
}

func TestLandsOnFreeLocationFacingGivenDirection(t *testing.T) {
	planetSize, _ := size.Square(2)
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{smallRock.In(*absoluteCoordinate.From(1, 2))})
	coordinate := absoluteCoordinate.From(1, 1)
	direction := North{}

	testRover := godModRover.LandFacing(direction, *coordinate, testPlanet)

	assert.Equal(t, *coordinate, testRover.Position())
	assert.Equal(t, direction, testRover.Direction())
}

func TestCanLandOnObstacle(t *testing.T) {
	planetSize, _ := size.Square(2)
	coordinate := absoluteCoordinate.From(1, 1)
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{smallRock.In(*coordinate)})

	testRover := godModRover.Land(*coordinate, testPlanet)

	assert.NotNil(t, testRover)
	assert.Equal(t, *coordinate, testRover.Position())
}
