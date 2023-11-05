package rover_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	rockyPlanet "mars_rover/internal/domain/planet/rocky_planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovesForward(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})

	testCases := []struct {
		name               string
		initialDirection   direction.Direction
		expectedCoordinate *coordinate.AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			initialDirection:   &direction.North{},
			expectedCoordinate: coordinate.NewAbsolute(5, 6),
		},
		{
			name:               "facing east",
			initialDirection:   &direction.East{},
			expectedCoordinate: coordinate.NewAbsolute(6, 5),
		},
		{
			name:               "facing south",
			initialDirection:   &direction.South{},
			expectedCoordinate: coordinate.NewAbsolute(5, 4),
		},
		{
			name:               "facing west",
			initialDirection:   &direction.West{},
			expectedCoordinate: coordinate.NewAbsolute(4, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			landingLocation, _ := location.From(*coordinate.NewAbsolute(5, 5), testCase.initialDirection)
			testRover := rover.Land(*landingLocation, testPlanetWithoutObstacles)

			testRover.MoveForward()

			expectedLocation, _ := location.From(*testCase.expectedCoordinate, testCase.initialDirection)
			assert.Equal(t, expectedLocation, testRover.Location())
		})
	}
}

func TestMovesBackward(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})

	testCases := []struct {
		name               string
		initialDirection   direction.Direction
		expectedCoordinate *coordinate.AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			initialDirection:   &direction.North{},
			expectedCoordinate: coordinate.NewAbsolute(5, 4),
		},
		{
			name:               "facing east",
			initialDirection:   &direction.East{},
			expectedCoordinate: coordinate.NewAbsolute(4, 5),
		},
		{
			name:               "facing south",
			initialDirection:   &direction.South{},
			expectedCoordinate: coordinate.NewAbsolute(5, 6),
		},
		{
			name:               "facing west",
			initialDirection:   &direction.West{},
			expectedCoordinate: coordinate.NewAbsolute(6, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			landingLocation, _ := location.From(*coordinate.NewAbsolute(5, 5), testCase.initialDirection)
			testRover := rover.Land(*landingLocation, testPlanetWithoutObstacles)

			testRover.MoveBackward()

			expectedLocation, _ := location.From(*testCase.expectedCoordinate, testCase.initialDirection)
			assert.Equal(t, expectedLocation, testRover.Location())
		})
	}
}

func TestTurnsRight(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	coord := *coordinate.NewAbsolute(5, 5)

	testCases := []struct {
		name              string
		initialDirection  direction.Direction
		expectedDirection direction.Direction
	}{
		{
			name:              "facing north",
			initialDirection:  &direction.North{},
			expectedDirection: &direction.East{},
		},
		{
			name:              "facing east",
			initialDirection:  &direction.East{},
			expectedDirection: &direction.South{},
		},
		{
			name:              "facing south",
			initialDirection:  &direction.South{},
			expectedDirection: &direction.West{},
		},
		{
			name:              "facing west",
			initialDirection:  &direction.West{},
			expectedDirection: &direction.North{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			landingLocation, _ := location.From(coord, testCase.initialDirection)
			testRover := rover.Land(*landingLocation, testPlanetWithoutObstacles)

			testRover.TurnRight()

			expectedLocation, _ := location.From(coord, testCase.expectedDirection)
			assert.Equal(t, expectedLocation, testRover.Location())
		})
	}
}

func TestTurnsLeft(t *testing.T) {
	planetSize, _ := size.From(10, 10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	coord := *coordinate.NewAbsolute(5, 5)

	testCases := []struct {
		name              string
		initialDirection  direction.Direction
		expectedDirection direction.Direction
	}{
		{
			name:              "facing north",
			initialDirection:  &direction.North{},
			expectedDirection: &direction.West{},
		},
		{
			name:              "facing east",
			initialDirection:  &direction.East{},
			expectedDirection: &direction.North{},
		},
		{
			name:              "facing south",
			initialDirection:  &direction.South{},
			expectedDirection: &direction.East{},
		},
		{
			name:              "facing west",
			initialDirection:  &direction.West{},
			expectedDirection: &direction.South{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			landingLocation, _ := location.From(*coordinate.NewAbsolute(5, 5), testCase.initialDirection)
			testRover := rover.Land(*landingLocation, testPlanetWithoutObstacles)

			testRover.TurnLeft()

			expectedLocation, _ := location.From(coord, testCase.expectedDirection)
			assert.Equal(t, expectedLocation, testRover.Location())
		})
	}
}
