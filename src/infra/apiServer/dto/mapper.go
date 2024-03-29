package dto

import (
	"fmt"
	. "mars_rover/src/action/mover"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/coordinates"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/planetMap"
)

func FromMovementResults(movementResults []MovementResult) MovementResponseDTO {
	var responseDTO MovementResponseDTO
	for _, movementResult := range movementResults {
		issue := issueFrom(movementResult)
		responseDTO.Results = append(responseDTO.Results, movementDTOFrom(issue, movementResult))
	}
	return responseDTO
}

func issueFrom(result MovementResult) string {
	if result.IssueDetected {
		return fmt.Sprintf("unable to move on command '%v'.", result.Cmd.String())
	}
	return ""
}

func movementDTOFrom(issue string, result MovementResult) SingleMovementDTO {
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

func planetDTOFromPlanet(planet Planet) PlanetDTO {
	size := planet.Size()
	return PlanetDTO{
		Width:     size.Width(),
		Height:    size.Height(),
		Obstacles: mapDomainToDTOObstacles(planet.Obstacles()),
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

func mapDomainToDTOCoordinates(coord Coordinates) []CoordinateDTO {
	var coordinatesDTO []CoordinateDTO
	for _, coordinate := range coord.List() {
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

func FromGame(game *Game) GameDTO {
	return GameDTO{
		Rover:  roverDTOFrom(game.Rover),
		Planet: planetDTOFromPlanet(game.Planet),
	}
}
