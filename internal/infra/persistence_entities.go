package infra

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

type RockyPlanetPersistenceEntity struct {
	Color     string                      `json:"color"`
	Size      SizePersistenceEntity       `json:"size"`
	Obstacles []ObstaclePersistenceEntity `json:"obstacles"`
}
