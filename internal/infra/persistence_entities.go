package infra

import (
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
	Coordinate CoordinatePersistenceEntity `json:"coordinate"`
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

func (r *SQLiteRepository) mapToPersistenceEntity(rover rover.Rover) RoverPersistenceEntity {
	obs := make([]ObstaclePersistenceEntity, len(rover.Map().Obstacles()))
	for i, o := range rover.Map().Obstacles() {
		coordinates := o.Coordinates()
		obs[i] = ObstaclePersistenceEntity{
			Coordinate: CoordinatePersistenceEntity{
				X: coordinates.X(),
				Y: coordinates.Y(),
			},
		}
	}

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
			Obstacles: obs,
		},
	}
}
