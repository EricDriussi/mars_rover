package dto

import (
	"fmt"
	. "mars_rover/src/action/move"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/planetMap"
)

func FromMovementResult(movementResult []MovementResult) MovementResponseDTO {
	var responseDTO MovementResponseDTO
	for _, result := range movementResult {
		issue := issueFrom(result)
		responseDTO.Results = append(responseDTO.Results, singleMovementDTOFrom(issue, result))
	}
	return responseDTO
}

func issueFrom(result MovementResult) string {
	if result.IssueDetected {
		return fmt.Sprintf("unable to move on command '%v'.", result.Cmd.String())
	}
	return ""
}

func singleMovementDTOFrom(issue string, result MovementResult) SingleMovementDTO {
	return SingleMovementDTO{
		Issue:      issue,
		Coordinate: coordinateDTOFrom(result.Coord),
		Direction:  result.Dir.CardinalPoint(),
	}
}

func FromDomainRover(rover Rover) CreateResponseDTO {
	planetMap := rover.Map()
	return CreateResponseDTO{
		Rover:  roverDTOFrom(rover),
		Planet: planetDTOFrom(planetMap),
	}
}

func planetDTOFrom(planetMap Map) PlanetDTO {
	return PlanetDTO{
		Width:     planetMap.Width(),
		Height:    planetMap.Height(),
		Obstacles: mapDomainToDTOObstacles(planetMap.Obstacles()),
	}
}

func roverDTOFrom(rover Rover) RoverDTO {
	coordinate := rover.Coordinate()
	return RoverDTO{
		Id:         rover.Id().String(),
		Coordinate: coordinateDTOFrom(coordinate),
		Direction:  rover.Direction().CardinalPoint(),
	}
}

func mapDomainToDTOObstacles(obstacles Obstacles) []ObstacleDTO {
	var obstaclesDTO []ObstacleDTO
	for _, obstacle := range obstacles.List() {
		coordinates := obstacle.Coordinates()
		obstaclesDTO = append(obstaclesDTO, mapDomainToDTOCoordinates(coordinates))
	}
	return obstaclesDTO
}

func mapDomainToDTOCoordinates(coord []AbsoluteCoordinate) []CoordinateDTO {
	var coordinatesDTO []CoordinateDTO
	for _, coordinate := range coord {
		coordinatesDTO = append(coordinatesDTO, coordinateDTOFrom(coordinate))
	}
	return coordinatesDTO
}

func coordinateDTOFrom(coordinate AbsoluteCoordinate) CoordinateDTO {
	return CoordinateDTO{
		X: coordinate.X(),
		Y: coordinate.Y(),
	}
}
