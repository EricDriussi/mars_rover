package controllers_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/godModRover"
	"mars_rover/src/domain/size"
	"mars_rover/src/infra/apiServer/controllers"
	"mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestReturnsRoverDTO(t *testing.T) {
	planetSize, _ := size.Square(2)
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []obstacle.Obstacle{})
	coordinate := absoluteCoordinate.From(1, 1)
	direction := North{}
	testRover := godModRover.LandFacing(uuid.New(), direction, *coordinate, testPlanet)
	mockAction := new(mocks.MockAction)
	mockAction.On("Random").Return(testRover, nil)

	dto, err := controllers.RandomGame(mockAction)

	assert.Nil(t, err)
	assert.Equal(t, dto.Rover.Id, testRover.Id().String())
	assert.Equal(t, dto.Rover.Direction, testRover.Direction().CardinalPoint())
	expectedCoordinate := testRover.Coordinate()
	assert.Equal(t, dto.Rover.Coordinate.X, expectedCoordinate.X())
	assert.Equal(t, dto.Rover.Coordinate.Y, expectedCoordinate.Y())
	expectedPlanet := testPlanet.Size()
	assert.Equal(t, dto.Planet.Width, expectedPlanet.Width())
	assert.Equal(t, dto.Planet.Height, expectedPlanet.Height())
	expectedObstacles := testPlanet.Obstacles()
	assert.Equal(t, len(dto.Planet.Obstacles), len(expectedObstacles.List()))
}

func TestReturnsErrorWhenActionFails(t *testing.T) {
	mockAction := new(mocks.MockAction)
	mockAction.On("Random").Return(new(mocks.MockRover), errors.New("test error"))

	_, err := controllers.RandomGame(mockAction)
	assert.NotNil(t, err)
}
