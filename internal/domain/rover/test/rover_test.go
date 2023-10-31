package rover_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovesForward(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})

	testCases := []struct {
		name               string
		initialDirection   direction.Direction
		expectedCoordinate coordinate.Coordinate
	}{
		{
			name:               "facing north",
			initialDirection:   &direction.North{},
			expectedCoordinate: *coordinate.New(5, 6),
		},
		{
			name:               "facing east",
			initialDirection:   &direction.East{},
			expectedCoordinate: *coordinate.New(6, 5),
		},
		{
			name:               "facing south",
			initialDirection:   &direction.South{},
			expectedCoordinate: *coordinate.New(5, 4),
		},
		{
			name:               "facing west",
			initialDirection:   &direction.West{},
			expectedCoordinate: *coordinate.New(4, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			landingLocation, _ := location.From(*coordinate.New(5, 5), testCase.initialDirection)
			testRover := rover.Land(*landingLocation, *testPlanetWithoutObstacles)

			testRover.MoveForward()

			expectedLocation, _ := location.From(testCase.expectedCoordinate, testCase.initialDirection)
			assert.True(t, expectedLocation.Equals(*testRover.Location()))
		})
	}
}

func TestMovesBackward(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})

	testCases := []struct {
		name               string
		initialDirection   direction.Direction
		expectedCoordinate coordinate.Coordinate
	}{
		{
			name:               "facing north",
			initialDirection:   &direction.North{},
			expectedCoordinate: *coordinate.New(5, 4),
		},
		{
			name:               "facing east",
			initialDirection:   &direction.East{},
			expectedCoordinate: *coordinate.New(4, 5),
		},
		{
			name:               "facing south",
			initialDirection:   &direction.South{},
			expectedCoordinate: *coordinate.New(5, 6),
		},
		{
			name:               "facing west",
			initialDirection:   &direction.West{},
			expectedCoordinate: *coordinate.New(6, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			landingLocation, _ := location.From(*coordinate.New(5, 5), testCase.initialDirection)
			testRover := rover.Land(*landingLocation, *testPlanetWithoutObstacles)

			testRover.MoveBackward()

			expectedLocation, _ := location.From(testCase.expectedCoordinate, testCase.initialDirection)
			assert.True(t, expectedLocation.Equals(*testRover.Location()))
		})
	}
}

func TestTurnsRight(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})

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
			landingLocation, _ := location.From(*coordinate.New(5, 5), testCase.initialDirection)
			testRover := rover.Land(*landingLocation, *testPlanetWithoutObstacles)

			testRover.TurnRight()

			assert.Equal(t, testCase.expectedFacing, testRover.Location().Orientation())
		})
	}
}

func TestTurnsLeft(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	testPlanetWithoutObstacles, _ := planet.Create(*planetSize, []obstacle.Obstacle{})

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
			landingLocation, _ := location.From(*coordinate.New(5, 5), testCase.initialDirection)
			testRover := rover.Land(*landingLocation, *testPlanetWithoutObstacles)

			testRover.TurnLeft()

			assert.Equal(t, testCase.expectedFacing, testRover.Location().Orientation())
		})
	}
}
