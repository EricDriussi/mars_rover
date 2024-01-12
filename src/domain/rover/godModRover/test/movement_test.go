package godModRover_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/godModRover"
	"mars_rover/src/domain/rover/id"
	. "mars_rover/src/test_helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovesForward(t *testing.T) {
	planet := SetupEmptyTestPlanetOfSize(t, 10)
	testCases := []struct {
		name               string
		initialDirection   Direction
		expectedCoordinate *AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			initialDirection:   &North{},
			expectedCoordinate: absoluteCoordinate.Build(5, 6),
		},
		{
			name:               "facing east",
			initialDirection:   &East{},
			expectedCoordinate: absoluteCoordinate.Build(6, 5),
		},
		{
			name:               "facing south",
			initialDirection:   &South{},
			expectedCoordinate: absoluteCoordinate.Build(5, 4),
		},
		{
			name:               "facing west",
			initialDirection:   &West{},
			expectedCoordinate: absoluteCoordinate.Build(4, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover := godModRover.LandFacing(id.New(), testCase.initialDirection, *absoluteCoordinate.Build(5, 5), planet)

			err := testRover.MoveForward()

			assert.Equal(t, *testCase.expectedCoordinate, testRover.Coordinate())
			assert.Nil(t, err)
		})
	}
}

func TestMovesBackward(t *testing.T) {
	planet := SetupEmptyTestPlanetOfSize(t, 10)
	testCases := []struct {
		name               string
		initialDirection   Direction
		expectedCoordinate *AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			initialDirection:   &North{},
			expectedCoordinate: absoluteCoordinate.Build(5, 4),
		},
		{
			name:               "facing east",
			initialDirection:   &East{},
			expectedCoordinate: absoluteCoordinate.Build(4, 5),
		},
		{
			name:               "facing south",
			initialDirection:   &South{},
			expectedCoordinate: absoluteCoordinate.Build(5, 6),
		},
		{
			name:               "facing west",
			initialDirection:   &West{},
			expectedCoordinate: absoluteCoordinate.Build(6, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover := godModRover.LandFacing(id.New(), testCase.initialDirection, *absoluteCoordinate.Build(5, 5), planet)

			err := testRover.MoveBackward()

			assert.Equal(t, *testCase.expectedCoordinate, testRover.Coordinate())
			assert.Nil(t, err)
		})
	}
}

func TestTurnsRight(t *testing.T) {
	planet := SetupEmptyTestPlanetOfSize(t, 10)
	coordinate := absoluteCoordinate.Build(5, 5)
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
			testRover := godModRover.LandFacing(id.New(), testCase.initialDirection, *coordinate, planet)

			testRover.TurnRight()

			assert.Equal(t, *coordinate, testRover.Coordinate())
			assert.Equal(t, testCase.expectedDirection, testRover.Direction())
		})
	}
}

func TestTurnsLeft(t *testing.T) {
	planet := SetupEmptyTestPlanetOfSize(t, 10)
	coordinate := absoluteCoordinate.Build(5, 5)
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
			testRover := godModRover.LandFacing(id.New(), testCase.initialDirection, *coordinate, planet)

			testRover.TurnLeft()

			assert.Equal(t, *coordinate, testRover.Coordinate())
			assert.Equal(t, testCase.expectedDirection, testRover.Direction())
		})
	}
}
