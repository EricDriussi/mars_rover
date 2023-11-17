package rover_test

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	smallRock "mars_rover/internal/domain/obstacle/small_rock"
	rockyPlanet "mars_rover/internal/domain/planet/rocky_planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLandsOnFreeLocation(t *testing.T) {
	planetSize, _ := size.Square(2)
	testPlanet, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{smallRock.In(*absoluteCoordinate.From(1, 2))})
	landingLocation, _ := location.From(*absoluteCoordinate.From(1, 1), &direction.North{})

	testRover, err := rover.Land(*landingLocation, testPlanet)

	assert.Nil(t, err)
	assert.Equal(t, landingLocation, testRover.Location())
}

func TestCannotLandOnObstacle(t *testing.T) {
	planetSize, _ := size.Square(2)
	testPlanet, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{smallRock.In(*absoluteCoordinate.From(1, 1))})
	landingLocation, _ := location.From(*absoluteCoordinate.From(1, 1), &direction.North{})

	testRover, err := rover.Land(*landingLocation, testPlanet)

	assert.Error(t, err)
	assert.Nil(t, testRover)
}

func TestMovesForward(t *testing.T) {
	planetSize, _ := size.Square(10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})

	testCases := []struct {
		name               string
		initialDirection   direction.Direction
		expectedCoordinate *absoluteCoordinate.AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			initialDirection:   &direction.North{},
			expectedCoordinate: absoluteCoordinate.From(5, 6),
		},
		{
			name:               "facing east",
			initialDirection:   &direction.East{},
			expectedCoordinate: absoluteCoordinate.From(6, 5),
		},
		{
			name:               "facing south",
			initialDirection:   &direction.South{},
			expectedCoordinate: absoluteCoordinate.From(5, 4),
		},
		{
			name:               "facing west",
			initialDirection:   &direction.West{},
			expectedCoordinate: absoluteCoordinate.From(4, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			landingLocation, _ := location.From(*absoluteCoordinate.From(5, 5), testCase.initialDirection)
			testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

			err := testRover.MoveForward()

			expectedLocation, _ := location.From(*testCase.expectedCoordinate, testCase.initialDirection)
			assert.Equal(t, expectedLocation, testRover.Location())
			assert.Nil(t, err)
		})
	}
}

func TestMovesBackward(t *testing.T) {
	planetSize, _ := size.Square(10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})

	testCases := []struct {
		name               string
		initialDirection   direction.Direction
		expectedCoordinate *absoluteCoordinate.AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			initialDirection:   &direction.North{},
			expectedCoordinate: absoluteCoordinate.From(5, 4),
		},
		{
			name:               "facing east",
			initialDirection:   &direction.East{},
			expectedCoordinate: absoluteCoordinate.From(4, 5),
		},
		{
			name:               "facing south",
			initialDirection:   &direction.South{},
			expectedCoordinate: absoluteCoordinate.From(5, 6),
		},
		{
			name:               "facing west",
			initialDirection:   &direction.West{},
			expectedCoordinate: absoluteCoordinate.From(6, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			landingLocation, _ := location.From(*absoluteCoordinate.From(5, 5), testCase.initialDirection)
			testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

			err := testRover.MoveBackward()

			expectedLocation, _ := location.From(*testCase.expectedCoordinate, testCase.initialDirection)
			assert.Equal(t, expectedLocation, testRover.Location())
			assert.Nil(t, err)
		})
	}
}

func TestTurnsRight(t *testing.T) {
	planetSize, _ := size.Square(10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	coord := *absoluteCoordinate.From(5, 5)

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
			testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

			testRover.TurnRight()

			expectedLocation, _ := location.From(coord, testCase.expectedDirection)
			assert.Equal(t, expectedLocation, testRover.Location())
		})
	}
}

func TestTurnsLeft(t *testing.T) {
	planetSize, _ := size.Square(10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create(*planetSize, []obstacle.Obstacle{})
	coord := *absoluteCoordinate.From(5, 5)

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
			landingLocation, _ := location.From(*absoluteCoordinate.From(5, 5), testCase.initialDirection)
			testRover, _ := rover.Land(*landingLocation, testPlanetWithoutObstacles)

			testRover.TurnLeft()

			expectedLocation, _ := location.From(coord, testCase.expectedDirection)
			assert.Equal(t, expectedLocation, testRover.Location())
		})
	}
}
