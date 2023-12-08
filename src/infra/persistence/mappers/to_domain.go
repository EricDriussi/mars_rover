package mappers

import (
	"errors"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/bigRock"
	. "mars_rover/src/domain/obstacle/obstacles"
	obstaclesModule "mars_rover/src/domain/obstacle/obstacles"
	"mars_rover/src/domain/obstacle/smallRock"
	. "mars_rover/src/domain/planet"
	"mars_rover/src/domain/planet/emptyPlanet"
	"mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/godModRover"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	s "mars_rover/src/domain/size"
	. "mars_rover/src/infra/persistence/entities"
)

func MapToDomainRover(roverEntity RoverEntity, planet Planet) (Rover, error) {
	direction, err := directionFromString(roverEntity.Direction)
	if err != nil {
		return nil, err
	}
	coordinate := absoluteCoordinate.From(roverEntity.Coordinate.X, roverEntity.Coordinate.Y)

	var rover Rover
	if roverEntity.GodMod {
		rover = godModRover.LandFacing(roverEntity.ID, direction, *coordinate, planet)
	} else {
		rover, err = wrappingCollidingRover.LandFacing(direction, *coordinate, planet)
		if err != nil {
			return nil, err
		}
	}
	return rover, nil
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
