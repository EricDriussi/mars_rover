package rover_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/direction"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapsLookingNorthMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingPosition, _ := coordinate.From(2, 0)

	testRover := rover.Land(*landingPosition, &direction.North{}, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedPosition, _ := coordinate.From(2, 3)
	assert.True(t, expectedPosition.Equals(testRover.Position()))
}

func TestWrapsLookingEastMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingPosition, _ := coordinate.From(0, 2)

	testRover := rover.Land(*landingPosition, &direction.East{}, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedPosition, _ := coordinate.From(3, 2)
	assert.True(t, expectedPosition.Equals(testRover.Position()))
}

func TestWrapsLookingSouthMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingPosition, _ := coordinate.From(2, 3)

	testRover := rover.Land(*landingPosition, &direction.South{}, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedPosition, _ := coordinate.From(2, 0)
	assert.True(t, expectedPosition.Equals(testRover.Position()))
}

func TestWrapsLookingWestMovingBackward(t *testing.T) {
	planetSize, _ := size.From(3, 3)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingPosition, _ := coordinate.From(3, 2)

	testRover := rover.Land(*landingPosition, &direction.West{}, *testPlanetWithoutObstacles)

	testRover.MoveBackward()

	expectedPosition, _ := coordinate.From(0, 2)
	assert.True(t, expectedPosition.Equals(testRover.Position()))
}
