package positionCalculator_test

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/direction"
	. "mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet/rockyPlanet"
	"mars_rover/internal/domain/rover/planetMap"
	"mars_rover/internal/domain/rover/wrappingCollidingRover/positionCalculator"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapsMovingForward(t *testing.T) {
	planetSize, _ := size.Square(3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{})
	testMap := planetMap.Of(testPlanetWithoutObstacles)

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
			futurePosition := positionCalculator.Forward(testCase.direction, *testCase.initialCoordinate, *testMap)
			assert.Equal(t, *testCase.expectedCoordinate, futurePosition)
		})
	}
}

func TestWrapsMovingBackwards(t *testing.T) {
	planetSize, _ := size.Square(3)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{})
	testMap := planetMap.Of(testPlanetWithoutObstacles)

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
			futurePosition := positionCalculator.Backward(testCase.direction, *testCase.initialCoordinate, *testMap)
			assert.Equal(t, *testCase.expectedCoordinate, futurePosition)
		})
	}
}
