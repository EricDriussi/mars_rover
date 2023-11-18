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

func TestMovesForward(t *testing.T) {
	planetSize, _ := size.Square(10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{})
	testMap := planetMap.Of(testPlanetWithoutObstacles)
	initialCoordinate := absoluteCoordinate.From(5, 5)

	testCases := []struct {
		name               string
		direction          Direction
		expectedCoordinate *AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			direction:          &North{},
			expectedCoordinate: absoluteCoordinate.From(5, 6),
		},
		{
			name:               "facing east",
			direction:          &East{},
			expectedCoordinate: absoluteCoordinate.From(6, 5),
		},
		{
			name:               "facing south",
			direction:          &South{},
			expectedCoordinate: absoluteCoordinate.From(5, 4),
		},
		{
			name:               "facing west",
			direction:          &West{},
			expectedCoordinate: absoluteCoordinate.From(4, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			futurePosition := positionCalculator.Forward(testCase.direction, *initialCoordinate, *testMap)
			assert.Equal(t, *testCase.expectedCoordinate, futurePosition)
		})
	}
}

func TestMovesBackward(t *testing.T) {
	planetSize, _ := size.Square(10)
	testPlanetWithoutObstacles, _ := rockyPlanet.Create("testColor", *planetSize, []Obstacle{})
	testMap := planetMap.Of(testPlanetWithoutObstacles)
	initialCoordinate := absoluteCoordinate.From(5, 5)

	testCases := []struct {
		name               string
		direction          Direction
		expectedCoordinate *AbsoluteCoordinate
	}{
		{
			name:               "facing north",
			direction:          &North{},
			expectedCoordinate: absoluteCoordinate.From(5, 4),
		},
		{
			name:               "facing east",
			direction:          &East{},
			expectedCoordinate: absoluteCoordinate.From(4, 5),
		},
		{
			name:               "facing south",
			direction:          &South{},
			expectedCoordinate: absoluteCoordinate.From(5, 6),
		},
		{
			name:               "facing west",
			direction:          &West{},
			expectedCoordinate: absoluteCoordinate.From(6, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			futurePosition := positionCalculator.Backward(testCase.direction, *initialCoordinate, *testMap)
			assert.Equal(t, *testCase.expectedCoordinate, futurePosition)
		})
	}
}
