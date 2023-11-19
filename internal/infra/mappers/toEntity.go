package mappers

import (
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle/obstacles"
	. "mars_rover/internal/domain/planet/rockyPlanet"
	. "mars_rover/internal/domain/rover/wrappingCollidingRover"
	. "mars_rover/internal/infra/entities"
)

func MapToPersistenceRover(rover WrappingCollidingRover) RoverPersistenceEntity {
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
func MapToPersistenceRockyPlanet(planet RockyPlanet) RockyPlanetPersistenceEntity {
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
