package controllers_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "mars_rover/src/action"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
	. "mars_rover/src/infra/apiServer/controllers"
	. "mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/test_helpers"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestBuildsAMovementResponseDTOWithRoverAndNoErrors(t *testing.T) {
	aRover := aRover()
	mockAction := new(MockAction)
	mockAction.On("MoveSequence").Return(MovementResult{
		Rover:          aRover,
		MovementErrors: &MovementErrors{},
		Error:          nil,
	}, nil)

	dto, err := MoveRover(mockAction, aMoveRequest())

	assert.Nil(t, err)
	assert.Empty(t, dto.Errors)
	// TODO.LM: actually moving the rover is not the responsibility of the controller
	// It's not relevant for this test, we just check that the response dto is built as expected
	assertRoversEqual(t, dto.Rover, aRover)
}

func TestBuildsAMovementResponseDTOWithRoverAndMovementErrors(t *testing.T) {
	aRover := aRover()
	movementErrors := &MovementErrors{}
	anError := "an error"
	anotherError := "another error"
	movementErrors.Add("aCommand", errors.New(anError))
	movementErrors.Add("anotherCommand", errors.New(anotherError))
	mockAction := new(MockAction)
	mockAction.On("MoveSequence").Return(MovementResult{
		Rover:          aRover,
		MovementErrors: movementErrors,
		Error:          nil,
	}, nil)

	dto, err := MoveRover(mockAction, aMoveRequest())

	assert.Nil(t, err)
	assert.Len(t, dto.Errors, 2)
	AssertContains(t, dto.Errors, anError)
	AssertContains(t, dto.Errors, anotherError)
	// TODO.LM: actually moving the rover is not the responsibility of the controller
	// It's not relevant for this test, we just check that the response dto is built as expected
	assertRoversEqual(t, dto.Rover, aRover)
}

func TestErrorsIfActionDoesNotSucceed(t *testing.T) {
	errMsg := "an error message"
	mockAction := new(MockAction)
	mockAction.On("MoveSequence").Return(MovementResult{
		Rover:          nil,
		MovementErrors: &MovementErrors{},
		Error:          errors.New(errMsg),
	}, nil)

	_, err := MoveRover(mockAction, aMoveRequest())

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), errMsg)
}

func aRover() Rover {
	aRover := new(MockRover)
	aRover.On("Id").Return(uuid.New())
	aRover.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	aRover.On("Direction").Return(North{})
	return aRover
}

func aMoveRequest() MoveRequest {
	return MoveRequest{
		Commands: "whatever",
		Id:       uuid.New().String(),
	}
}

func assertRoversEqual(t *testing.T, roverDTO RoverDTO, rover Rover) {
	assert.Equal(t, roverDTO.Id, rover.Id().String())
	assert.Equal(t, roverDTO.Direction, rover.Direction().CardinalPoint())
	expectedCoordinate := rover.Coordinate()
	assert.Equal(t, roverDTO.Coordinate.X, expectedCoordinate.X())
	assert.Equal(t, roverDTO.Coordinate.Y, expectedCoordinate.Y())
}
