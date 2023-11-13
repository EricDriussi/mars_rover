package infra_test

import (
	"errors"
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/infra"
)

func mapToDomainRover(roverData infra.RoverPersistenceEntity, planet planet.Planet) (rover.Rover, error) {
	dir, err := directionFromString(roverData.Location.Direction)
	if err != nil {
		return nil, err
	}

	loc, err := location.From(
		*coordinate.NewAbsolute(roverData.Location.CurrentCoord.X, roverData.Location.CurrentCoord.Y),
		dir,
	)
	if err != nil {
		return nil, err
	}

	roverInstance, err := rover.Land(*loc, planet)
	if err != nil {
		return nil, err
	}

	return roverInstance, nil
}

func directionFromString(dirStr string) (direction.Direction, error) {
	switch dirStr {
	case "N":
		return &direction.North{}, nil
	case "S":
		return &direction.South{}, nil
	case "E":
		return &direction.East{}, nil
	case "W":
		return &direction.West{}, nil
	}
	return nil, errors.New("Invalid direction")
}
