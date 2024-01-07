package test_helpers

import (
	"errors"
	"github.com/stretchr/testify/assert"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
	obs "mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/obstacles"
	"mars_rover/src/domain/planet"
	"mars_rover/src/domain/size"
	"testing"
)

func SuccessfulRoverFunc() func() error {
	return func() error {
		return nil
	}
}

func RoverFunc(fn func() error) func() error {
	return func() error {
		return fn()
	}
}

func FailedRoverFunc() func() error {
	return func() error {
		return errors.New("an error")
	}
}

func SetupEmptyTestPlanetOfSize(t *testing.T, n int) planet.Planet {
	planetSize, err := size.Square(n)
	assert.Nil(t, err)
	testPlanet, err := planet.CreatePlanet("testColor", *planetSize, *obstacles.Empty())
	assert.Nil(t, err)
	return testPlanet
}

func SetupPlanetOfSizeWithObstacleIn(t *testing.T, n int, coordinate AbsoluteCoordinate) planet.Planet {
	planetSize, err := size.Square(n)
	assert.Nil(t, err)
	coord, err := coordinates.New(coordinate)
	assert.Nil(t, err)
	obstacle, err := obs.CreateObstacle(*coord)
	assert.Nil(t, err)
	obss, err := obstacles.FromList(obstacle)
	assert.Nil(t, err)
	testPlanet, err := planet.CreatePlanet("testColor", *planetSize, *obss)
	assert.Nil(t, err)
	return testPlanet
}
