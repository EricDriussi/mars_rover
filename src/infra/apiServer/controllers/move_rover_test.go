package controllers_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "mars_rover/src/action"
	. "mars_rover/src/action/command"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
	. "mars_rover/src/infra/apiServer/controllers"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestBuildsAMovementResponseDTOWhenNoMovementIssuesAreReported(t *testing.T) {
	resultWithNoIssues := MovementResult{
		Cmd:           Forward,
		IssueDetected: false,
		Coord:         *absoluteCoordinate.From(1, 1),
		Dir:           North{},
	}
	mockAction := new(MockAction)
	mockAction.On("MoveSequence").Return([]MovementResult{resultWithNoIssues}, nil)

	movementResponseDTO, err := MoveRover(mockAction, aMoveRequest())

	assert.Nil(t, err)
	singleMovementDTO := movementResponseDTO.Results[0]
	assert.Empty(t, singleMovementDTO.Issue)
	assert.Equal(t, singleMovementDTO.Coordinate.X, resultWithNoIssues.Coord.X())
	assert.Equal(t, singleMovementDTO.Coordinate.Y, resultWithNoIssues.Coord.Y())
	assert.Equal(t, singleMovementDTO.Direction, resultWithNoIssues.Dir.CardinalPoint())
}

func TestBuildsAMovementResponseDTOWhenMovementIssuesAreReported(t *testing.T) {
	resultWithIssues := MovementResult{
		Cmd:           Forward,
		IssueDetected: true,
		Coord:         *absoluteCoordinate.From(1, 1),
		Dir:           North{},
	}
	mockAction := new(MockAction)
	mockAction.On("MoveSequence").Return([]MovementResult{resultWithIssues}, nil)

	movementResponseDTO, err := MoveRover(mockAction, aMoveRequest())

	assert.Nil(t, err)
	singleMovementDTO := movementResponseDTO.Results[0]
	assert.Contains(t, singleMovementDTO.Issue, resultWithIssues.Cmd.ToString())
	assert.Equal(t, singleMovementDTO.Coordinate.X, resultWithIssues.Coord.X())
	assert.Equal(t, singleMovementDTO.Coordinate.Y, resultWithIssues.Coord.Y())
	assert.Equal(t, singleMovementDTO.Direction, resultWithIssues.Dir.CardinalPoint())
}

func TestBuildsAMovementResponseDTOReportingMovementIssuesIfFound(t *testing.T) {
	resultWithNoIssues := MovementResult{
		Cmd:           Forward,
		IssueDetected: false,
		Coord:         *absoluteCoordinate.From(1, 1),
		Dir:           North{},
	}
	resultWithIssues := MovementResult{
		Cmd:           Forward,
		IssueDetected: true,
		Coord:         *absoluteCoordinate.From(1, 1),
		Dir:           North{},
	}
	mockAction := new(MockAction)
	mockAction.On("MoveSequence").Return([]MovementResult{resultWithNoIssues, resultWithIssues}, nil)

	movementResponseDTO, err := MoveRover(mockAction, aMoveRequest())

	assert.Nil(t, err)
	noIssuesMovementDTO := movementResponseDTO.Results[0]
	assert.Empty(t, noIssuesMovementDTO.Issue)
	movementDTOWithIssue := movementResponseDTO.Results[1]
	assert.Contains(t, movementDTOWithIssue.Issue, resultWithIssues.Cmd.ToString())
}

func TestErrorsIfActionDoesNotSucceed(t *testing.T) {
	mockAction := new(MockAction)
	errMsg := "an error message"
	mockAction.On("MoveSequence").Return([]MovementResult{}, errors.New(errMsg))

	_, err := MoveRover(mockAction, aMoveRequest())

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), errMsg)
}

func aMoveRequest() MoveRequest {
	return MoveRequest{
		Commands: "whatever",
		Id:       uuid.New().String(),
	}
}
