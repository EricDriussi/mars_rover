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

func beforeEach() (*planet.Planet, *coordinate.Coordinate) {
	planetSize, _ := size.From(10, 10)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})
	landingPosition, _ := coordinate.From(5, 5)

	return testPlanetWithoutObstacles, landingPosition
}

func TestMovesForward(t *testing.T) {
	testPlanetWithoutObstacles, landingPosition := beforeEach()

	testCases := []struct {
		name             string
		initialDirection direction.Direction
		expectedX        int
		expectedY        int
	}{
		{
			name:             "facing north",
			initialDirection: &direction.North{},
			expectedX:        5,
			expectedY:        6,
		},
		{
			name:             "facing east",
			initialDirection: &direction.East{},
			expectedX:        6,
			expectedY:        5,
		},
		{
			name:             "facing south",
			initialDirection: &direction.South{},
			expectedX:        5,
			expectedY:        4,
		},
		{
			name:             "facing west",
			initialDirection: &direction.West{},
			expectedX:        4,
			expectedY:        5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover := rover.Land(*landingPosition, testCase.initialDirection, *testPlanetWithoutObstacles)

			testRover.MoveForward()

			expectedPosition, _ := coordinate.From(testCase.expectedX, testCase.expectedY)
			assert.True(t, expectedPosition.Equals(testRover.Position()))
		})
	}
}

func TestMovesBackward(t *testing.T) {
	testPlanetWithoutObstacles, landingPosition := beforeEach()

	testCases := []struct {
		name             string
		initialDirection direction.Direction
		expectedX        int
		expectedY        int
	}{
		{
			name:             "facing north",
			initialDirection: &direction.North{},
			expectedX:        5,
			expectedY:        4,
		},
		{
			name:             "facing east",
			initialDirection: &direction.East{},
			expectedX:        4,
			expectedY:        5,
		},
		{
			name:             "facing south",
			initialDirection: &direction.South{},
			expectedX:        5,
			expectedY:        6,
		},
		{
			name:             "facing west",
			initialDirection: &direction.West{},
			expectedX:        6,
			expectedY:        5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover := rover.Land(*landingPosition, testCase.initialDirection, *testPlanetWithoutObstacles)

			testRover.MoveBackward()

			expectedPosition, _ := coordinate.From(testCase.expectedX, testCase.expectedY)
			assert.True(t, expectedPosition.Equals(testRover.Position()))
		})
	}
}

func TestTurnsRight(t *testing.T) {
	testPlanetWithoutObstacles, landingPosition := beforeEach()

	testCases := []struct {
		name             string
		initialDirection direction.Direction
		expectedFacing   string
	}{
		{
			name:             "facing north",
			initialDirection: &direction.North{},
			expectedFacing:   "E",
		},
		{
			name:             "facing east",
			initialDirection: &direction.East{},
			expectedFacing:   "S",
		},
		{
			name:             "facing south",
			initialDirection: &direction.South{},
			expectedFacing:   "W",
		},
		{
			name:             "facing west",
			initialDirection: &direction.West{},
			expectedFacing:   "N",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover := rover.Land(*landingPosition, testCase.initialDirection, *testPlanetWithoutObstacles)

			testRover.TurnRight()

			assert.Equal(t, testCase.expectedFacing, testRover.Direction().CardinalPoint())
		})
	}
}

func TestTurnsLeft(t *testing.T) {
	testPlanetWithoutObstacles, landingPosition := beforeEach()

	testCases := []struct {
		name             string
		initialDirection direction.Direction
		expectedFacing   string
	}{
		{
			name:             "facing north",
			initialDirection: &direction.North{},
			expectedFacing:   "W",
		},
		{
			name:             "facing east",
			initialDirection: &direction.East{},
			expectedFacing:   "N",
		},
		{
			name:             "facing south",
			initialDirection: &direction.South{},
			expectedFacing:   "E",
		},
		{
			name:             "facing west",
			initialDirection: &direction.West{},
			expectedFacing:   "S",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover := rover.Land(*landingPosition, testCase.initialDirection, *testPlanetWithoutObstacles)

			testRover.TurnLeft()

			assert.Equal(t, testCase.expectedFacing, testRover.Direction().CardinalPoint())
		})
	}
}
