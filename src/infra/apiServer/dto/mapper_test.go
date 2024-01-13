package dto_test

import (
	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action/mover"
	. "mars_rover/src/action/mover/command"
	. "mars_rover/src/domain"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
	"mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/obstacle/obstacles"
	"mars_rover/src/domain/planet"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/id"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	. "mars_rover/src/domain/rover/wrappingCollidingRover"
	"mars_rover/src/domain/size"
	"mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/infra/apiServer/dto"
	"mars_rover/src/test_helpers/mocks"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestBuildsAMovementResponseDTOFromMovementResultWithNoMovementIssues(t *testing.T) {
	resultWithIssues := MovementResult{
		Cmd:           Forward,
		IssueDetected: false,
		Coord:         *absoluteCoordinate.Build(1, 1),
		Dir:           North{},
	}

	movementResponseDTO := dto.FromMovementResults([]MovementResult{resultWithIssues})

	singleMovementDTO := movementResponseDTO.Results[0]
	assert.Empty(t, singleMovementDTO.Issue)
	assertMoveDTOContainsDataFrom(t, singleMovementDTO, resultWithIssues)
}

func TestBuildsAMovementResponseDTOFromMovementResultWithMovementIssues(t *testing.T) {
	resultWithIssues := MovementResult{
		Cmd:           Forward,
		IssueDetected: true,
		Coord:         *absoluteCoordinate.Build(1, 1),
		Dir:           North{},
	}

	movementResponseDTO := dto.FromMovementResults([]MovementResult{resultWithIssues})

	singleMovementDTO := movementResponseDTO.Results[0]
	assert.Contains(t, singleMovementDTO.Issue, resultWithIssues.Cmd.String())
	assertMoveDTOContainsDataFrom(t, singleMovementDTO, resultWithIssues)
}

func assertMoveDTOContainsDataFrom(t *testing.T, singleMovementDTO SingleMovementDTO, resultWithIssues MovementResult) {
	assert.Equal(t, singleMovementDTO.Coordinate.X, resultWithIssues.Coord.X())
	assert.Equal(t, singleMovementDTO.Coordinate.Y, resultWithIssues.Coord.Y())
	assert.Equal(t, singleMovementDTO.Direction, resultWithIssues.Dir.CardinalPoint())
}

func TestBuildsAMovementResponseDTOFromMultipleMovementResults(t *testing.T) {
	resultWithNoIssues := MovementResult{
		Cmd:           Forward,
		IssueDetected: false,
		Coord:         *absoluteCoordinate.Build(1, 1),
		Dir:           North{},
	}
	resultWithIssues := MovementResult{
		Cmd:           Forward,
		IssueDetected: true,
		Coord:         *absoluteCoordinate.Build(1, 1),
		Dir:           North{},
	}

	movementResponseDTO := dto.FromMovementResults([]MovementResult{resultWithNoIssues, resultWithIssues})

	noIssuesMovementDTO := movementResponseDTO.Results[0]
	assert.Empty(t, noIssuesMovementDTO.Issue)
	movementDTOWithIssue := movementResponseDTO.Results[1]
	assert.Contains(t, movementDTOWithIssue.Issue, resultWithIssues.Cmd.String())
}

func TestBuildsACreateResponseDTOFromARover(t *testing.T) {
	mockPlanet := mocks.PlanetWithNoObstaclesOfSize(t, 10)
	testRover, err := wrappingCollidingRover.LandFacing(id.New(), North{}, *absoluteCoordinate.Build(1, 1), mockPlanet)
	assert.Nil(t, err)

	createResponseDTO := dto.FromDomainRover(testRover)

	assertCreateDTOContainsRoverData(t, createResponseDTO, testRover)
	assertCreateDTOContainsPlanetData(t, createResponseDTO, mockPlanet)
}

func TestBuildsAGameDTOFromAGame(t *testing.T) {
	sizeLimit, err := size.Square(10)
	assert.Nil(t, err)
	mockObstacle := new(MockObstacle)
	mockObstacle.On("Occupies", Anything).Return(false)
	mockObstacle.On("IsBeyond", Anything).Return(false)
	coords, err := coordinates.New(*absoluteCoordinate.Build(1, 1))
	assert.Nil(t, err)
	mockObstacle.On("Coordinates").Return(*coords)
	obs, err := obstacles.FromList(mockObstacle)
	assert.Nil(t, err)
	testPlanet, err := planet.CreatePlanet("testColor", *sizeLimit, *obs)
	assert.Nil(t, err)
	testRover, err := wrappingCollidingRover.LandFacing(id.New(), North{}, *absoluteCoordinate.Build(1, 1), testPlanet)
	assert.Nil(t, err)
	testGame := &Game{
		Rover:  testRover,
		Planet: testPlanet,
	}

	loadResponseDTO := dto.FromGame(testGame)

	assertGameDTOContainsRoverData(t, loadResponseDTO, testRover)
	assertGameDTOContainsPlanetData(t, loadResponseDTO, testPlanet)
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

func assertGameDTOContainsRoverData(t *testing.T, gameDTO GameDTO, testRover *WrappingCollidingRover) {
	assert.Equal(t, gameDTO.Rover.Id, testRover.Id().String())
	testCoordinate := testRover.Coordinate()
	assert.Equal(t, gameDTO.Rover.Coordinate.X, testCoordinate.X())
	assert.Equal(t, gameDTO.Rover.Coordinate.Y, testCoordinate.Y())
	assert.Equal(t, gameDTO.Rover.Direction, testRover.Direction().CardinalPoint())
	testMap := testRover.Map()
	assert.Equal(t, gameDTO.Planet.Width, testMap.Width())
	assert.Equal(t, gameDTO.Planet.Height, testMap.Height())
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

func assertGameDTOContainsPlanetData(t *testing.T, gameDTO GameDTO, planet Planet) {
	planetSize := planet.Size()
	planetDTO := gameDTO.Planet
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
		coords := obs.Coordinates()
		for j, coord := range coords.List() {
			assert.Equal(t, coordinateDTOS[j].X, coord.X())
			assert.Equal(t, coordinateDTOS[j].Y, coord.Y())
		}
	}
}
