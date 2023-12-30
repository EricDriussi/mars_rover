package wrappingCollidingRover_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/uuid"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	"mars_rover/src/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapsMovingForward(t *testing.T) {
	planetSize, _ := size.Square(4)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{})
	testCases := []struct {
		name               string
		direction          Direction
		initialCoordinate  *AbsoluteCoordinate
		expectedCoordinate *AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			direction:          &North{},
			initialCoordinate:  absoluteCoordinate.From(2, 3),
			expectedCoordinate: absoluteCoordinate.From(2, 0),
		},
		{
			name:               "facing east",
			direction:          &East{},
			initialCoordinate:  absoluteCoordinate.From(3, 2),
			expectedCoordinate: absoluteCoordinate.From(0, 2),
		},
		{
			name:               "facing south",
			direction:          &South{},
			initialCoordinate:  absoluteCoordinate.From(2, 0),
			expectedCoordinate: absoluteCoordinate.From(2, 3),
		},
		{
			name:               "facing west",
			direction:          &West{},
			initialCoordinate:  absoluteCoordinate.From(0, 2),
			expectedCoordinate: absoluteCoordinate.From(3, 2),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover, _ := wrappingCollidingRover.LandFacing(uuid.New(), testCase.direction, *testCase.initialCoordinate, testPlanetWithoutObstacles)

			err := testRover.MoveForward()

			assert.Nil(t, err)
			assert.Equal(t, *testCase.expectedCoordinate, testRover.Coordinate())
		})
	}
}

func TestWrapsMovingBackwards(t *testing.T) {
	planetSize, _ := size.Square(4)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{})
	testCases := []struct {
		name               string
		direction          Direction
		initialCoordinate  *AbsoluteCoordinate
		expectedCoordinate *AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			direction:          &North{},
			initialCoordinate:  absoluteCoordinate.From(2, 0),
			expectedCoordinate: absoluteCoordinate.From(2, 3),
		},
		{
			name:               "facing east",
			direction:          &East{},
			initialCoordinate:  absoluteCoordinate.From(0, 2),
			expectedCoordinate: absoluteCoordinate.From(3, 2),
		},
		{
			name:               "facing south",
			direction:          &South{},
			initialCoordinate:  absoluteCoordinate.From(2, 3),
			expectedCoordinate: absoluteCoordinate.From(2, 0),
		},
		{
			name:               "facing west",
			direction:          &West{},
			initialCoordinate:  absoluteCoordinate.From(3, 2),
			expectedCoordinate: absoluteCoordinate.From(0, 2),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover, _ := wrappingCollidingRover.LandFacing(uuid.New(), testCase.direction, *testCase.initialCoordinate, testPlanetWithoutObstacles)

			err := testRover.MoveBackward()

			assert.Nil(t, err)
			assert.Equal(t, *testCase.expectedCoordinate, testRover.Coordinate())
		})
	}
}
