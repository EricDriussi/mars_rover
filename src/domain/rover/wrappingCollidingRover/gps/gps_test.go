package gps_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/id"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	"mars_rover/src/domain/rover/wrappingCollidingRover/gps"
	. "mars_rover/src/test_helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatesAhead(t *testing.T) {
	planet := SetupEmptyTestPlanetOfSize(t, 10)
	coordinate := absoluteCoordinate.Build(5, 5)
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
			rover, err := wrappingCollidingRover.LandFacing(id.New(), testCase.initialDirection, *coordinate, planet)
			assert.Nil(t, err)
			GPS := gps.Bind(rover)

			calculatedCoordinates := GPS.Ahead()

			assert.Equal(t, *testCase.expectedCoordinate, calculatedCoordinates)
		})
	}
}

func TestCalculatesBehind(t *testing.T) {
	planet := SetupEmptyTestPlanetOfSize(t, 10)
	coordinate := absoluteCoordinate.Build(5, 5)
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
			rover, err := wrappingCollidingRover.LandFacing(id.New(), testCase.initialDirection, *coordinate, planet)
			assert.Nil(t, err)
			GPS := gps.Bind(rover)

			calculatedCoordinates := GPS.Behind()

			assert.Equal(t, *testCase.expectedCoordinate, calculatedCoordinates)
		})
	}
}

func TestCalculatesWrappingAhead(t *testing.T) {
	planet := SetupEmptyTestPlanetOfSize(t, 4)
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
			expectedCoordinate: absoluteCoordinate.Build(2, 0),
		},
		{
			name:               "facing east",
			direction:          &East{},
			initialCoordinate:  absoluteCoordinate.Build(3, 2),
			expectedCoordinate: absoluteCoordinate.Build(0, 2),
		},
		{
			name:               "facing south",
			direction:          &South{},
			initialCoordinate:  absoluteCoordinate.Build(2, 0),
			expectedCoordinate: absoluteCoordinate.Build(2, 3),
		},
		{
			name:               "facing west",
			direction:          &West{},
			initialCoordinate:  absoluteCoordinate.Build(0, 2),
			expectedCoordinate: absoluteCoordinate.Build(3, 2),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rover, err := wrappingCollidingRover.LandFacing(id.New(), testCase.direction, *testCase.initialCoordinate, planet)
			assert.Nil(t, err)
			GPS := gps.Bind(rover)

			calculatedCoordinates := GPS.Ahead()

			assert.Equal(t, *testCase.expectedCoordinate, calculatedCoordinates)
		})
	}
}

func TestCalculatesWrappingBehind(t *testing.T) {
	planet := SetupEmptyTestPlanetOfSize(t, 4)
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
			expectedCoordinate: absoluteCoordinate.Build(2, 3),
		},
		{
			name:               "facing east",
			direction:          &East{},
			initialCoordinate:  absoluteCoordinate.Build(0, 2),
			expectedCoordinate: absoluteCoordinate.Build(3, 2),
		},
		{
			name:               "facing south",
			direction:          &South{},
			initialCoordinate:  absoluteCoordinate.Build(2, 3),
			expectedCoordinate: absoluteCoordinate.Build(2, 0),
		},
		{
			name:               "facing west",
			direction:          &West{},
			initialCoordinate:  absoluteCoordinate.Build(3, 2),
			expectedCoordinate: absoluteCoordinate.Build(0, 2),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rover, err := wrappingCollidingRover.LandFacing(id.New(), testCase.direction, *testCase.initialCoordinate, planet)
			assert.Nil(t, err)
			GPS := gps.Bind(rover)

			calculatedCoordinates := GPS.Behind()

			assert.Equal(t, *testCase.expectedCoordinate, calculatedCoordinates)
		})
	}
}
