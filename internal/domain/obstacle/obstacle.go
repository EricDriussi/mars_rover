package obstacle

import (
	coord "mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/size"
)

type Obstacle struct {
	Coordinate *coord.Coordinate
}

func In(coordinate *coord.Coordinate) *Obstacle {
	return &Obstacle{coordinate}
}

func (this Obstacle) IsWithinLimit(size size.Size) bool {
	return this.Coordinate.IsWithin(size)
}
