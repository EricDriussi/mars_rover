package dto_test

import (
	"github.com/stretchr/testify/assert"
	. "mars_rover/src/action"
	. "mars_rover/src/action/command"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/infra/apiServer/dto"
	"testing"
)

func TestBuildsAMovementResponseDTOFromMovementResultWithNoMovementIssues(t *testing.T) {
	resultWithIssues := MovementResult{
		Cmd:           Forward,
		IssueDetected: false,
		Coord:         *absoluteCoordinate.From(1, 1),
		Dir:           North{},
	}

	movementResponseDTO := dto.FromActionResult([]MovementResult{resultWithIssues})

	singleMovementDTO := movementResponseDTO.Results[0]
	assert.Empty(t, singleMovementDTO.Issue)
	assert.Equal(t, singleMovementDTO.Coordinate.X, resultWithIssues.Coord.X())
	assert.Equal(t, singleMovementDTO.Coordinate.Y, resultWithIssues.Coord.Y())
	assert.Equal(t, singleMovementDTO.Direction, resultWithIssues.Dir.CardinalPoint())
}

func TestBuildsAMovementResponseDTOFromMovementResultWithMovementIssues(t *testing.T) {
	resultWithIssues := MovementResult{
		Cmd:           Forward,
		IssueDetected: true,
		Coord:         *absoluteCoordinate.From(1, 1),
		Dir:           North{},
	}

	movementResponseDTO := dto.FromActionResult([]MovementResult{resultWithIssues})

	singleMovementDTO := movementResponseDTO.Results[0]
	assert.Contains(t, singleMovementDTO.Issue, resultWithIssues.Cmd.ToString())
	assert.Equal(t, singleMovementDTO.Coordinate.X, resultWithIssues.Coord.X())
	assert.Equal(t, singleMovementDTO.Coordinate.Y, resultWithIssues.Coord.Y())
	assert.Equal(t, singleMovementDTO.Direction, resultWithIssues.Dir.CardinalPoint())
}

func TestBuildsAMovementResponseDTOFromMultipleMovementResults(t *testing.T) {
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

	movementResponseDTO := dto.FromActionResult([]MovementResult{resultWithNoIssues, resultWithIssues})

	noIssuesMovementDTO := movementResponseDTO.Results[0]
	assert.Empty(t, noIssuesMovementDTO.Issue)
	movementDTOWithIssue := movementResponseDTO.Results[1]
	assert.Contains(t, movementDTOWithIssue.Issue, resultWithIssues.Cmd.ToString())
}
