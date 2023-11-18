package infra

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/obstacle/obstacles"
	"mars_rover/internal/domain/rover"
)

type RoverPersistenceEntity struct {
	PlanetMap  PlanetMapPersistenceEntity  `json:"planetMap"`
	Coordinate CoordinatePersistenceEntity `json:"coordinate"`
	Direction  string                      `json:"direction"`
}

type PlanetMapPersistenceEntity struct {
	Size      SizePersistenceEntity       `json:"size"`
	Obstacles []ObstaclePersistenceEntity `json:"obstacles"`
}

type ObstaclePersistenceEntity struct {
	Coordinates []CoordinatePersistenceEntity `json:"coordinates"`
}

type CoordinatePersistenceEntity struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type SizePersistenceEntity struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (r *SQLiteRepository) mapToPersistenceRover(rover rover.Rover) RoverPersistenceEntity {
	currPosition := rover.Position()
	roverMap := rover.Map()
	return RoverPersistenceEntity{
		Coordinate: CoordinatePersistenceEntity{
			X: currPosition.X(),
			Y: currPosition.Y(),
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

func mapToPersistenceCoordinates(coordinates []absoluteCoordinate.AbsoluteCoordinate) []CoordinatePersistenceEntity {
	coords := make([]CoordinatePersistenceEntity, len(coordinates))
	for i, c := range coordinates {
		coords[i] = CoordinatePersistenceEntity{
			X: c.X(),
			Y: c.Y(),
		}
	}
	return coords
}

func mapToPersistenceObstacles(obstacles obstacles.Obstacles) []ObstaclePersistenceEntity {
	obs := make([]ObstaclePersistenceEntity, len(obstacles.List()))
	for i, o := range obstacles.List() {
		coordinates := o.Coordinates()
		obs[i] = ObstaclePersistenceEntity{
			Coordinates: mapToPersistenceCoordinates(coordinates),
		}
	}
	return obs
}
