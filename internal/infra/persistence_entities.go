package infra

import (
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
)

type RoverPersistenceEntity struct {
	Location  LocationPersistenceEntity  `json:"location"`
	PlanetMap PlanetMapPersistenceEntity `json:"planetMap"`
}

type PlanetMapPersistenceEntity struct {
	Size      size.Size                   `json:"size"`
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
			Size: size.Size{
				Width:  rover.Map().Size().Width,
				Height: rover.Map().Size().Height,
			},
			Obstacles: obs,
		},
	}
}
