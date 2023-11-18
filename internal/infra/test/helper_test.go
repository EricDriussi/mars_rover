package infra_test

import (
	"errors"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/direction"
	. "mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/infra"
)

func mapToDomainRover(roverData RoverPersistenceEntity, planet Planet) (Rover, error) {
	dir, err := directionFromString(roverData.Direction)
	if err != nil {
		return nil, err
	}

	coordinate := absoluteCoordinate.From(roverData.Coordinate.X, roverData.Coordinate.Y)

	roverInstance, err := rover.LandFacing(dir, *coordinate, planet)
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
