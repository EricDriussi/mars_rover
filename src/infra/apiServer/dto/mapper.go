package dto

import (
	"fmt"
	. "mars_rover/src/action"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/rover"
)

func FromMovementResult(movementResult []MovementResult) MovementResponseDTO {
	var responseDTO MovementResponseDTO
	for _, result := range movementResult {
		issue := ""
		if result.IssueDetected {
			issue = fmt.Sprintf("unable to move on command '%v'.", result.Cmd.String())
		}
		responseDTO.Results = append(responseDTO.Results, SingleMovementDTO{
			Issue: issue,
			Coordinate: CoordinateDTO{
				X: result.Coord.X(),
				Y: result.Coord.Y(),
			},
			Direction: result.Dir.CardinalPoint(),
		})
	}
	return responseDTO
}

func FromDomainRover(rover Rover) CreateResponseDTO {
	coordinate := rover.Coordinate()
	planetMap := rover.Map()
	return CreateResponseDTO{
		Rover: RoverDTO{
			Id: rover.Id().String(),
			Coordinate: CoordinateDTO{
				X: coordinate.X(),
				Y: coordinate.Y(),
			},
			Direction: rover.Direction().CardinalPoint(),
		},
		Planet: PlanetDTO{
			Width:     planetMap.Width(),
			Height:    planetMap.Height(),
			Obstacles: mapDomainToDTOObstacles(planetMap.Obstacles()),
		},
	}
}

func mapDomainToDTOObstacles(obstacles Obstacles) []ObstacleDTO {
	var obstaclesDTO []ObstacleDTO
	for _, obstacle := range obstacles.List() {
		coordinates := obstacle.Coordinates()
		obstaclesDTO = append(obstaclesDTO, ObstacleDTO{
			Coordinate: mapDomainToDTOCoordinates(coordinates),
		})
	}
	return obstaclesDTO
}

func mapDomainToDTOCoordinates(c []AbsoluteCoordinate) []CoordinateDTO {
	var coordinatesDTO []CoordinateDTO
	for _, coordinate := range c {
		coordinatesDTO = append(coordinatesDTO, CoordinateDTO{
			X: coordinate.X(),
			Y: coordinate.Y(),
		})
	}
	return coordinatesDTO
}
