package infra

import (
	"errors"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle/obstacles"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/domain/rover/direction"
	"mars_rover/internal/domain/rover/wrappingCollidingRover"
)

func mapToPersistenceRover(rover Rover) RoverPersistenceEntity {
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

func MapToDomainRover(roverData RoverPersistenceEntity, planet Planet) (Rover, error) {
	dir, err := directionFromString(roverData.Direction)
	if err != nil {
		return nil, err
	}

	coordinate := absoluteCoordinate.From(roverData.Coordinate.X, roverData.Coordinate.Y)

	roverInstance, err := wrappingCollidingRover.LandFacing(dir, *coordinate, planet)
	if err != nil {
		return nil, err
	}

	return roverInstance, nil
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
