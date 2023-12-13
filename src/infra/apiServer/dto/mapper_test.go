package dto_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action"
	. "mars_rover/src/action/command"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	"mars_rover/src/domain/size"
	"mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestBuildsAMovementResponseDTOFromMovementResultWithNoMovementIssues(t *testing.T) {
	resultWithIssues := MovementResult{
		Cmd:           Forward,
		IssueDetected: false,
		Coord:         *absoluteCoordinate.From(1, 1),
		Dir:           North{},
	}

	movementResponseDTO := dto.FromMovementResult([]MovementResult{resultWithIssues})

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

	movementResponseDTO := dto.FromMovementResult([]MovementResult{resultWithIssues})

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

	movementResponseDTO := dto.FromMovementResult([]MovementResult{resultWithNoIssues, resultWithIssues})

	noIssuesMovementDTO := movementResponseDTO.Results[0]
	assert.Empty(t, noIssuesMovementDTO.Issue)
	movementDTOWithIssue := movementResponseDTO.Results[1]
	assert.Contains(t, movementDTOWithIssue.Issue, resultWithIssues.Cmd.ToString())
}

func TestBuildsACreateResponseDTOFromARover(t *testing.T) {
	testSize, _ := size.Square(10)
	mockPlanet := new(MockPlanet)
	mockPlanet.On("Size").Return(*testSize)
	mockObstacle := new(MockObstacle)
	mockObstacle.On("Occupies", Anything).Return(false)
	mockObstacle.On("Coordinates").Return([]absoluteCoordinate.AbsoluteCoordinate{})
	testObstacles := obstacles.FromList([]obstacle.Obstacle{mockObstacle})
	mockPlanet.On("Obstacles").Return(*testObstacles)
	testRover, _ := wrappingCollidingRover.LandFacing(uuid.New(), North{}, *absoluteCoordinate.From(1, 1), mockPlanet)

	createResponseDTO := dto.FromDomainRover(testRover)

	assert.Equal(t, createResponseDTO.Rover.Id, testRover.Id().String())
	testCoordinate := testRover.Coordinate()
	assert.Equal(t, createResponseDTO.Rover.Coordinate.X, testCoordinate.X())
	assert.Equal(t, createResponseDTO.Rover.Coordinate.Y, testCoordinate.Y())
	assert.Equal(t, createResponseDTO.Rover.Direction, testRover.Direction().CardinalPoint())
	testMap := testRover.Map()
	assert.Equal(t, createResponseDTO.Planet.Width, testMap.Width())
	assert.Equal(t, createResponseDTO.Planet.Height, testMap.Height())
	assert.Len(t, createResponseDTO.Planet.Obstacles, len(testObstacles.List()))
}