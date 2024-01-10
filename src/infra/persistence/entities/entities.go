package entities

import (
	. "mars_rover/src/domain/rover/uuid"
)

type RoverEntity struct {
	ID         UUID             `json:"id"`
	PlanetMap  MapEntity        `json:"planetMap"`
	Coordinate CoordinateEntity `json:"coordinate"`
	Direction  string           `json:"direction"`
	Type       string           `json:"type"`
	PlanetId   int              `json:"planetId"`
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
