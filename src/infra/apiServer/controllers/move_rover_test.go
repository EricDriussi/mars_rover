package controllers_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "mars_rover/src/action"
	. "mars_rover/src/action/test"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/godModRover"
	"mars_rover/src/domain/size"
	. "mars_rover/src/infra/apiServer/controllers"
	"testing"
)

// TODO: more and better tests
func TestReturnsMovementResponseDTOWithUpdatedRover(t *testing.T) {
	planetSize, _ := size.Square(2)
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []obstacle.Obstacle{})
	coordinate := absoluteCoordinate.From(1, 1)
	direction := North{}
	testRover := godModRover.LandFacing(uuid.New(), direction, *coordinate, testPlanet)
	mockAction := new(MockAction)
	movementResult := MovementResult{
		Rover:          testRover,
		MovementErrors: &MovementErrors{},
		Error:          nil,
	}
	mockAction.On("MoveSequence").Return(movementResult, nil)
	moveRequest := MoveRequest{
		Commands: "f",
		Id:       testRover.Id().String(),
	}

	dto, err := MoveRover(mockAction, moveRequest)

	assert.Nil(t, err)
	assert.Equal(t, dto.Rover.Id, testRover.Id().String())
	assert.Equal(t, dto.Rover.Direction, testRover.Direction().CardinalPoint())
	expectedCoordinate := testRover.Coordinate()
	assert.Equal(t, dto.Rover.Coordinate.X, expectedCoordinate.X())
	assert.Equal(t, dto.Rover.Coordinate.Y, expectedCoordinate.Y())
	assert.Equal(t, len(dto.Errors), 0)
}

func TestReturnsErrorWhenMoveActionFails(t *testing.T) {
	mockAction := new(MockAction)
	errMsg := "test error"
	movementResult := MovementResult{
		Rover:          nil,
		MovementErrors: &MovementErrors{},
		Error:          errors.New(errMsg),
	}
	mockAction.On("MoveSequence").Return(movementResult, nil)
	moveRequest := MoveRequest{
		Commands: "f",
		Id:       uuid.New().String(),
	}

	_, err := MoveRover(mockAction, moveRequest)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), errMsg)
}
