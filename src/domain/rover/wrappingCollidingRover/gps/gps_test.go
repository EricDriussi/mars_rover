package gps_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	"mars_rover/src/domain/rover/wrappingCollidingRover/gps"
	"mars_rover/src/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatesAhead(t *testing.T) {
	planetSize, _ := size.Square(10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{})

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
			testRover, _ := wrappingCollidingRover.LandFacing(testCase.initialDirection, *coordinate, testPlanetWithoutObstacles)

			GPS := gps.Bind(testRover)
			calculatedCoordinates := GPS.Ahead()

			assert.Equal(t, *testCase.expectedCoordinate, calculatedCoordinates)
		})
	}
}

func TestCalculatesBehind(t *testing.T) {
	planetSize, _ := size.Square(10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{})

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
			testRover, _ := wrappingCollidingRover.LandFacing(testCase.initialDirection, *coordinate, testPlanetWithoutObstacles)

			GPS := gps.Bind(testRover)
			calculatedCoordinates := GPS.Behind()

			assert.Equal(t, *testCase.expectedCoordinate, calculatedCoordinates)
		})
	}
}

func TestCalculatesWrappingAhead(t *testing.T) {
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
			testRover, _ := wrappingCollidingRover.LandFacing(
				testCase.direction,
				*testCase.initialCoordinate,
				testPlanetWithoutObstacles,
			)

			GPS := gps.Bind(testRover)
			calculatedCoordinates := GPS.Ahead()

			assert.Equal(t, *testCase.expectedCoordinate, calculatedCoordinates)
		})
	}
}

func TestCalculatesWrappingBehind(t *testing.T) {
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
			testRover, _ := wrappingCollidingRover.LandFacing(
				testCase.direction,
				*testCase.initialCoordinate,
				testPlanetWithoutObstacles,
			)

			GPS := gps.Bind(testRover)
			calculatedCoordinates := GPS.Behind()

			assert.Equal(t, *testCase.expectedCoordinate, calculatedCoordinates)
		})
	}
}
