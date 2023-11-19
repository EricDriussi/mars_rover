package entities

type RoverEntity struct {
	PlanetMap  MapEntity        `json:"planetMap"`
	Coordinate CoordinateEntity `json:"coordinate"`
	Direction  string           `json:"direction"`
}

type MapEntity struct {
	Size      SizeEntity       `json:"size"`
	Obstacles []ObstacleEntity `json:"obstacles"`
}

type ObstacleEntity struct {
	Coordinates []CoordinateEntity `json:"coordinates"`
}

type CoordinateEntity struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type SizeEntity struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type PlanetEntity struct {
	Color     string           `json:"color"`
	Size      SizeEntity       `json:"size"`
	Obstacles []ObstacleEntity `json:"obstacles"`
}
