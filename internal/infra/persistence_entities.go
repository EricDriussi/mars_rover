package infra

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/obstacle/obstacles"
	"mars_rover/internal/domain/rover"
)

type RoverPersistenceEntity struct {
	Location  LocationPersistenceEntity  `json:"location"`
	PlanetMap PlanetMapPersistenceEntity `json:"planetMap"`
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

type LocationPersistenceEntity struct {
	CurrentCoord CoordinatePersistenceEntity `json:"currentCoord"`
	FutureCoord  CoordinatePersistenceEntity `json:"futureCoord"`
	Direction    string                      `json:"direction"`
}

type SizePersistenceEntity struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (r *SQLiteRepository) mapToPersistenceRover(rover rover.Rover) RoverPersistenceEntity {
	currPosition := rover.Location().Position()
	futurePosition := rover.Location().WillBeAt()
	size := rover.Map().Size()
	return RoverPersistenceEntity{
		Location: LocationPersistenceEntity{
			CurrentCoord: CoordinatePersistenceEntity{
				X: currPosition.X(),
				Y: currPosition.Y(),
			},
			FutureCoord: CoordinatePersistenceEntity{
				X: futurePosition.X(),
				Y: futurePosition.Y(),
			},
			Direction: rover.Location().Direction().CardinalPoint(),
		},
		PlanetMap: PlanetMapPersistenceEntity{
			Size: SizePersistenceEntity{
				Width:  size.Width(),
				Height: size.Height(),
			},
			Obstacles: mapToPersistenceObstacles(rover.Map().Obstacles()),
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
