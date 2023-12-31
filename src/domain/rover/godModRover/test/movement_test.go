package godModRover_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
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

// TODO: review domain tests...
func TestMovesForward(t *testing.T) {
	planetSize, _ := size.Square(10)
	rock := smallRock.In(*absoluteCoordinate.From(1, 1))
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{&rock})

	testCases := []struct {
		name               string
		initialDirection   Direction
		expectedCoordinate *AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			initialDirection:   &North{},
			expectedCoordinate: absoluteCoordinate.From(5, 6),
		},
		{
			name:               "facing east",
			initialDirection:   &East{},
			expectedCoordinate: absoluteCoordinate.From(6, 5),
		},
		{
			name:               "facing south",
			initialDirection:   &South{},
			expectedCoordinate: absoluteCoordinate.From(5, 4),
		},
		{
			name:               "facing west",
			initialDirection:   &West{},
			expectedCoordinate: absoluteCoordinate.From(4, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			coordinate := absoluteCoordinate.From(5, 5)
			testRover := godModRover.LandFacing(uuid.New(), testCase.initialDirection, *coordinate, testPlanet)

			err := testRover.MoveForward()

			assert.Equal(t, *testCase.expectedCoordinate, testRover.Coordinate())
			assert.Nil(t, err)
		})
	}
}

func TestMovesBackward(t *testing.T) {
	planetSize, _ := size.Square(10)
	rock := smallRock.In(*absoluteCoordinate.From(1, 1))
	// TODO: all these "A, _ := ..." should be using assert.Nil on the error
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{&rock})

	testCases := []struct {
		name               string
		initialDirection   Direction
		expectedCoordinate *AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			initialDirection:   &North{},
			expectedCoordinate: absoluteCoordinate.From(5, 4),
		},
		{
			name:               "facing east",
			initialDirection:   &East{},
			expectedCoordinate: absoluteCoordinate.From(4, 5),
		},
		{
			name:               "facing south",
			initialDirection:   &South{},
			expectedCoordinate: absoluteCoordinate.From(5, 6),
		},
		{
			name:               "facing west",
			initialDirection:   &West{},
			expectedCoordinate: absoluteCoordinate.From(6, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			coordinate := absoluteCoordinate.From(5, 5)
			testRover := godModRover.LandFacing(uuid.New(), testCase.initialDirection, *coordinate, testPlanet)

			err := testRover.MoveBackward()

			assert.Equal(t, *testCase.expectedCoordinate, testRover.Coordinate())
			assert.Nil(t, err)
		})
	}
}

func TestTurnsRight(t *testing.T) {
	planetSize, _ := size.Square(10)
	rock := smallRock.In(*absoluteCoordinate.From(1, 1))
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{&rock})
	coord := absoluteCoordinate.From(5, 5)

	testCases := []struct {
		name              string
		initialDirection  Direction
		expectedDirection Direction
	}{
		{
			name:              "facing north",
			initialDirection:  &North{},
			expectedDirection: &East{},
		},
		{
			name:              "facing east",
			initialDirection:  &East{},
			expectedDirection: &South{},
		},
		{
			name:              "facing south",
			initialDirection:  &South{},
			expectedDirection: &West{},
		},
		{
			name:              "facing west",
			initialDirection:  &West{},
			expectedDirection: &North{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover := godModRover.LandFacing(uuid.New(), testCase.initialDirection, *coord, testPlanet)

			testRover.TurnRight()

			assert.Equal(t, *coord, testRover.Coordinate())
			assert.Equal(t, testCase.expectedDirection, testRover.Direction())
		})
	}
}

func TestTurnsLeft(t *testing.T) {
	planetSize, _ := size.Square(10)
	rock := smallRock.In(*absoluteCoordinate.From(1, 1))
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{&rock})
	coord := absoluteCoordinate.From(5, 5)

	testCases := []struct {
		name              string
		initialDirection  Direction
		expectedDirection Direction
	}{
		{
			name:              "facing north",
			initialDirection:  &North{},
			expectedDirection: &West{},
		},
		{
			name:              "facing east",
			initialDirection:  &East{},
			expectedDirection: &North{},
		},
		{
			name:              "facing south",
			initialDirection:  &South{},
			expectedDirection: &East{},
		},
		{
			name:              "facing west",
			initialDirection:  &West{},
			expectedDirection: &South{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover := godModRover.LandFacing(uuid.New(), testCase.initialDirection, *coord, testPlanet)

			testRover.TurnLeft()

			assert.Equal(t, *coord, testRover.Coordinate())
			assert.Equal(t, testCase.expectedDirection, testRover.Direction())
		})
	}
}
