package mappers

import (
	"errors"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/bigRock"
	. "mars_rover/internal/domain/obstacle/obstacles"
	obstaclesModule "mars_rover/internal/domain/obstacle/obstacles"
	"mars_rover/internal/domain/obstacle/smallRock"
	. "mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/planet/emptyPlanet"
	"mars_rover/internal/domain/planet/rockyPlanet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/domain/rover/direction"
	"mars_rover/internal/domain/rover/godModRover"
	. "mars_rover/internal/domain/rover/planetMap"
	pm "mars_rover/internal/domain/rover/planetMap"
	"mars_rover/internal/domain/rover/wrappingCollidingRover"
	s "mars_rover/internal/domain/size"
	. "mars_rover/internal/infra/entities"
)

func MapToDomainRovers(roverEntities []RoverEntity) ([]Rover, error) {
	rovers := make([]Rover, 0, len(roverEntities))
	for _, roverEntity := range roverEntities {
		rover, err := MapToDomainRover(roverEntity)
		if err != nil {
			return nil, err
		}
		rovers = append(rovers, rover)
	}
	return rovers, nil

}

func MapToDomainRover(roverEntity RoverEntity) (Rover, error) {
	direction, err := directionFromString(roverEntity.Direction)
	if err != nil {
		return nil, err
	}
	coordinate := absoluteCoordinate.From(roverEntity.Coordinate.X, roverEntity.Coordinate.Y)

	domainMap, err := mapToDomainMap(roverEntity.PlanetMap)
	if err != nil {
		return nil, err
	}

	if roverEntity.GodMod {
		return godModRover.Start(direction, *coordinate, domainMap), nil
	}
	return wrappingCollidingRover.Start(direction, *coordinate, domainMap)
}

func mapToDomainMap(planetMap MapEntity) (Map, error) {
	size, err := s.Shape(planetMap.Size.Width, planetMap.Size.Height)
	if err != nil {
		return Map{}, err
	}
	obstacles := mapToDomainObstacles(planetMap.Obstacles)
	return *pm.Of(*size, obstacles), nil
}

func directionFromString(dirStr string) (Direction, error) {
	switch dirStr {
	case "N":
		return &North{}, nil
	case "S":
		return &South{}, nil
	case "E":
		return &East{}, nil
	case "W":
		return &West{}, nil
	}
	return nil, errors.New("Invalid direction")
}

func MapToDomainPlanets(planetEntities []PlanetEntity) ([]Planet, error) {
	planets := make([]Planet, 0, len(planetEntities))
	for _, planetEntity := range planetEntities {
		planet, err := MapToDomainPlanet(planetEntity)
		if err != nil {
			return nil, err
		}
		planets = append(planets, planet)
	}
	return planets, nil
}

func MapToDomainPlanet(planetEntity PlanetEntity) (Planet, error) {
	color := planetEntity.Color
	size, err := s.Square(planetEntity.Size.Width)
	if err != nil {
		return nil, err
	}
	obstacles := mapToDomainObstacles(planetEntity.Obstacles)
	if len(obstacles.List()) == 0 {
		return emptyPlanet.Create(color, *size)
	}
	return rockyPlanet.Create(color, *size, obstacles.List())
}

func mapToDomainObstacles(obstacles []ObstacleEntity) Obstacles {
	list := make([]Obstacle, 0, len(obstacles))
	for _, obstacle := range obstacles {
		coordinates := mapToDomainCoordinates(obstacle.Coordinates)
		if len(coordinates) <= 1 {
			rock := smallRock.In(coordinates[0])
			list = append(list, &rock)
		} else {
			rock := bigRock.In(coordinates)
			list = append(list, &rock)
		}
	}
	return *obstaclesModule.FromList(list)
}

func mapToDomainCoordinates(coordinates []CoordinateEntity) []AbsoluteCoordinate {
	list := make([]AbsoluteCoordinate, 0, len(coordinates))
	for _, coordinate := range coordinates {
		list = append(list, *absoluteCoordinate.From(coordinate.X, coordinate.Y))
	}
	return list
}
