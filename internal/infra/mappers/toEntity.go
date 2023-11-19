package mappers

import (
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle/obstacles"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/domain/size"
	. "mars_rover/internal/infra/entities"
)

func MapToPersistenceRover(rover Rover) RoverEntity {
	coordinate := rover.Coordinate()
	roverMap := rover.Map()
	return RoverEntity{
		Coordinate: mapToPersistenceCoordinate(coordinate),
		Direction:  rover.Direction().CardinalPoint(),
		PlanetMap: MapEntity{
			Size: SizeEntity{
				Width:  roverMap.Width(),
				Height: roverMap.Height(),
			},
			Obstacles: mapToPersistenceObstacles(roverMap.Obstacles()),
		},
	}
}

func mapToPersistenceCoordinate(coordinate AbsoluteCoordinate) CoordinateEntity {
	return CoordinateEntity{
		X: coordinate.X(),
		Y: coordinate.Y(),
	}
}

func mapToPersistenceCoordinates(coordinates []AbsoluteCoordinate) []CoordinateEntity {
	coords := make([]CoordinateEntity, len(coordinates))
	for i, c := range coordinates {
		coords[i] = mapToPersistenceCoordinate(c)
	}
	return coords
}

func mapToPersistenceObstacles(obstacles Obstacles) []ObstacleEntity {
	obs := make([]ObstacleEntity, len(obstacles.List()))
	for i, o := range obstacles.List() {
		coordinates := o.Coordinates()
		obs[i] = ObstacleEntity{
			Coordinates: mapToPersistenceCoordinates(coordinates),
		}
	}
	return obs
}

func mapToPersistenceSize(size Size) SizeEntity {
	return SizeEntity{
		Width:  size.Width(),
		Height: size.Height(),
	}
}

func MapToPersistenceRockyPlanet(planet Planet) RockyPlanetEntity {
	size := planet.Size()
	return RockyPlanetEntity{
		Color:     planet.Color(),
		Size:      mapToPersistenceSize(size),
		Obstacles: mapToPersistenceObstacles(planet.Obstacles()),
	}
}
