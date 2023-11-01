package rock

import (
	coord "mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
)

type Rock struct {
	coordinate coord.Coordinate
}

func In(coordinate coord.Coordinate) obstacle.Obstacle {
	return &Rock{coordinate}
}

func (this Rock) Coordinate() coord.Coordinate {
	return this.coordinate
}

func (this Rock) IsBeyond(size size.Size) bool {
	return this.coordinate.IsOutsideOf(size)
}
