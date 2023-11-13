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
		obs[i] = ObstaclePersistenceEntity{
			Coordinate: CoordinatePersistenceEntity{
				X: o.Coordinates().X(),
				Y: o.Coordinates().Y(),
			},
		}
	}

	return RoverPersistenceEntity{
		Location: LocationPersistenceEntity{
			CurrentCoord: CoordinatePersistenceEntity{
				X: rover.Location().Position().X(),
				Y: rover.Location().Position().Y(),
			},
			FutureCoord: CoordinatePersistenceEntity{
				X: rover.Location().WillBeAt().X(),
				Y: rover.Location().WillBeAt().Y(),
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
