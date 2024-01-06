package godModRover_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/godModRover"
	"mars_rover/src/domain/rover/uuid"
	. "mars_rover/src/test_helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesNotWrapMovingForward(t *testing.T) {
	testPlanet := SetupEmptyTestPlanetOfSize(t, 3)
	testCases := []struct {
		name               string
		direction          Direction
		initialCoordinate  *AbsoluteCoordinate
		expectedCoordinate *AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			direction:          &North{},
			initialCoordinate:  absoluteCoordinate.Build(2, 3),
			expectedCoordinate: absoluteCoordinate.Build(2, 4),
		},
		{
			name:               "facing east",
			direction:          &East{},
			initialCoordinate:  absoluteCoordinate.Build(3, 2),
			expectedCoordinate: absoluteCoordinate.Build(4, 2),
		},
		{
			name:               "facing south",
			direction:          &South{},
			initialCoordinate:  absoluteCoordinate.Build(2, 0),
			expectedCoordinate: absoluteCoordinate.Build(2, -1),
		},
		{
			name:               "facing west",
			direction:          &West{},
			initialCoordinate:  absoluteCoordinate.Build(0, 2),
			expectedCoordinate: absoluteCoordinate.Build(-1, 2),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover := godModRover.LandFacing(uuid.New(), testCase.direction, *testCase.initialCoordinate, testPlanet)

			err := testRover.MoveForward()

			assert.Nil(t, err)
			assert.Equal(t, *testCase.expectedCoordinate, testRover.Coordinate())
		})
	}
}

func TestDoesNotWrapMovingBackwards(t *testing.T) {
	testPlanet := SetupEmptyTestPlanetOfSize(t, 3)
	testCases := []struct {
		name               string
		direction          Direction
		initialCoordinate  *AbsoluteCoordinate
		expectedCoordinate *AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			direction:          &North{},
			initialCoordinate:  absoluteCoordinate.Build(2, 0),
			expectedCoordinate: absoluteCoordinate.Build(2, -1),
		},
		{
			name:               "facing east",
			direction:          &East{},
			initialCoordinate:  absoluteCoordinate.Build(0, 2),
			expectedCoordinate: absoluteCoordinate.Build(-1, 2),
		},
		{
			name:               "facing south",
			direction:          &South{},
			initialCoordinate:  absoluteCoordinate.Build(2, 3),
			expectedCoordinate: absoluteCoordinate.Build(2, 4),
		},
		{
			name:               "facing west",
			direction:          &West{},
			initialCoordinate:  absoluteCoordinate.Build(3, 2),
			expectedCoordinate: absoluteCoordinate.Build(4, 2),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover := godModRover.LandFacing(uuid.New(), testCase.direction, *testCase.initialCoordinate, testPlanet)

			err := testRover.MoveBackward()

			assert.Nil(t, err)
			assert.Equal(t, *testCase.expectedCoordinate, testRover.Coordinate())
		})
	}
}
