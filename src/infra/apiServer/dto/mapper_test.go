package dto_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action/move"
	. "mars_rover/src/action/move/command"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	. "mars_rover/src/domain/rover/wrappingCollidingRover"
	"mars_rover/src/domain/size"
	"mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/infra/apiServer/dto"
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
	assertMoveDTOContains(t, singleMovementDTO, resultWithIssues)
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
	assert.Contains(t, singleMovementDTO.Issue, resultWithIssues.Cmd.String())
	assertMoveDTOContains(t, singleMovementDTO, resultWithIssues)
}

func assertMoveDTOContains(t *testing.T, singleMovementDTO SingleMovementDTO, resultWithIssues MovementResult) {
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
	assert.Contains(t, movementDTOWithIssue.Issue, resultWithIssues.Cmd.String())
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

	assertCreateDTOContainsRoverData(t, createResponseDTO, testRover)
	assertCreateDTOContainsPlanetData(t, createResponseDTO, mockPlanet)
}

func assertCreateDTOContainsRoverData(t *testing.T, createResponseDTO CreateResponseDTO, testRover *WrappingCollidingRover) {
	assert.Equal(t, createResponseDTO.Rover.Id, testRover.Id().String())
	testCoordinate := testRover.Coordinate()
	assert.Equal(t, createResponseDTO.Rover.Coordinate.X, testCoordinate.X())
	assert.Equal(t, createResponseDTO.Rover.Coordinate.Y, testCoordinate.Y())
	assert.Equal(t, createResponseDTO.Rover.Direction, testRover.Direction().CardinalPoint())
	testMap := testRover.Map()
	assert.Equal(t, createResponseDTO.Planet.Width, testMap.Width())
	assert.Equal(t, createResponseDTO.Planet.Height, testMap.Height())
}

func assertCreateDTOContainsPlanetData(t *testing.T, createResponseDTO CreateResponseDTO, planet Planet) {
	planetSize := planet.Size()
	planetDTO := createResponseDTO.Planet
	assert.Equal(t, planetDTO.Width, planetSize.Width())
	assert.Equal(t, planetDTO.Height, planetSize.Height())
	planetObstacles := planet.Obstacles()
	obstaclesDTO := planetDTO.Obstacles
	assert.Len(t, obstaclesDTO, len(planetObstacles.List()))
	assertSameCoordinates(t, obstaclesDTO, planetObstacles)
}

func assertSameCoordinates(t *testing.T, createResponseDTO []ObstacleDTO, obst Obstacles) {
	for i, obs := range obst.List() {
		coordinateDTOS := createResponseDTO[i]
		for j, coord := range obs.Coordinates() {
			assert.Equal(t, coordinateDTOS[j].X, coord.X())
			assert.Equal(t, coordinateDTOS[j].Y, coord.Y())
		}
	}
}
