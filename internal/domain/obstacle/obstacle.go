package obstacle

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/size"
)

type Obstacle struct {
	Position *coordinate.Coordinate
}

func In(position *coordinate.Coordinate) *Obstacle {
	return &Obstacle{position}
}

func (this Obstacle) IsWithinLimit(size size.Size) bool {
	return this.Position.IsWithin(size)
}
