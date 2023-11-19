package mappers

import (
	"errors"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle"
	. "mars_rover/internal/domain/obstacle/obstacles"
	obstaclesModule "mars_rover/internal/domain/obstacle/obstacles"
	"mars_rover/internal/domain/obstacle/smallRock"
	. "mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/planet/rockyPlanet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/domain/rover/direction"
	. "mars_rover/internal/domain/rover/wrappingCollidingRover"
	s "mars_rover/internal/domain/size"
	. "mars_rover/internal/infra/entities"
)

// TODO: This is cheating, re-thing mapping strategy so that
// if only one coord, then smallRock, else bigRock
// same should happen with planets and list of obstacles
func mapToDomainObstacles(obstacles []ObstaclePersistenceEntity) Obstacles {
	ob := make([]Obstacle, len(obstacles))
	for i, o := range obstacles {
		coordinates := mapToDomainCoordinates(o.Coordinates)
		r := smallRock.In(coordinates[0])
		ob[i] = &r
	}
	return *obstaclesModule.New(ob)
}

func mapToDomainCoordinates(coordinates []CoordinatePersistenceEntity) []AbsoluteCoordinate {
	coords := make([]AbsoluteCoordinate, len(coordinates))
	for i, c := range coordinates {
		coords[i] = *absoluteCoordinate.From(c.X, c.Y)
	}
	return coords
}

func MapToDomainRover(roverData RoverPersistenceEntity, planet Planet) (Rover, error) {
	dir, err := directionFromString(roverData.Direction)
	if err != nil {
		return nil, err
	}

	coordinate := absoluteCoordinate.From(roverData.Coordinate.X, roverData.Coordinate.Y)

	roverInstance, err := LandFacing(dir, *coordinate, planet)
	if err != nil {
		return nil, err
	}

	return roverInstance, nil
}

func MapToDomainPlanet(planetData RockyPlanetPersistenceEntity) (Planet, error) {
	color := planetData.Color
	s, err := s.Square(planetData.Size.Width)
	if err != nil {
		return nil, err
	}
	obstacles := mapToDomainObstacles(planetData.Obstacles)
	return rockyPlanet.Create(color, *s, obstacles.List())
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
