package obstacle

import (
	coord "mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/size"
)

type Obstacle interface {
	IsBeyond(size.Size) bool
	Coordinate() coord.Coordinate
}

type Rock struct {
	coordinate coord.Coordinate
}

func In(coordinate coord.Coordinate) Obstacle {
	return &Rock{coordinate}
}

func (this Rock) Coordinate() coord.Coordinate {
	return this.coordinate
}

func (this Rock) IsBeyond(size size.Size) bool {
	return this.coordinate.IsOutsideOf(size)
}
