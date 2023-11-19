package infra

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
	. "mars_rover/internal/domain/planet/rockyPlanet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/domain/rover/direction"
	. "mars_rover/internal/domain/rover/wrappingCollidingRover"
	s "mars_rover/internal/domain/size"
)

func mapToPersistenceRover(rover WrappingCollidingRover) RoverPersistenceEntity {
	coordinate := rover.Coordinate()
	roverMap := rover.Map()
	return RoverPersistenceEntity{
		Coordinate: CoordinatePersistenceEntity{
			X: coordinate.X(),
			Y: coordinate.Y(),
		},
		Direction: rover.Direction().CardinalPoint(),
		PlanetMap: PlanetMapPersistenceEntity{
			Size: SizePersistenceEntity{
				Width:  roverMap.Width(),
				Height: roverMap.Height(),
			},
			Obstacles: mapToPersistenceObstacles(roverMap.Obstacles()),
		},
	}
}

func mapToPersistenceCoordinates(coordinates []AbsoluteCoordinate) []CoordinatePersistenceEntity {
	coords := make([]CoordinatePersistenceEntity, len(coordinates))
	for i, c := range coordinates {
		coords[i] = CoordinatePersistenceEntity{
			X: c.X(),
			Y: c.Y(),
		}
	}
	return coords
}

func mapToPersistenceObstacles(obstacles Obstacles) []ObstaclePersistenceEntity {
	obs := make([]ObstaclePersistenceEntity, len(obstacles.List()))
	for i, o := range obstacles.List() {
		coordinates := o.Coordinates()
		obs[i] = ObstaclePersistenceEntity{
			Coordinates: mapToPersistenceCoordinates(coordinates),
		}
	}
	return obs
}

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

func mapToPersistenceRockyPlanet(planet RockyPlanet) RockyPlanetPersistenceEntity {
	size := planet.Size()
	return RockyPlanetPersistenceEntity{
		Color: planet.Color(),
		Size: SizePersistenceEntity{
			Width:  size.Width(),
			Height: size.Height(),
		},
		Obstacles: mapToPersistenceObstacles(planet.Obstacles()),
	}
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
