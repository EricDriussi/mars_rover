package controllers

import (
	. "mars_rover/src/action"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/infra/apiServer/dto"
)

func RandomGame(action Action) (CreateResponseDTO, error) {
	curiosity, err := action.Random()
	if err != nil {
		return CreateResponseDTO{}, err
	}

	coordinate := curiosity.Coordinate()
	planetMap := curiosity.Map()
	return CreateResponseDTO{
		Rover: RoverDTO{
			Id: curiosity.Id().String(),
			Coordinate: CoordinateDTO{
				X: coordinate.X(),
				Y: coordinate.Y(),
			},
			Direction: curiosity.Direction().CardinalPoint(),
		},
		Planet: PlanetDTO{
			Width:     planetMap.Width(),
			Height:    planetMap.Height(),
			Obstacles: mapDomainToDTOObstacles(planetMap.Obstacles()),
		},
	}, nil
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
