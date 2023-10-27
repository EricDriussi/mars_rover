package obstacle

import (
	"mars_rover/internal/domain/position"
	"mars_rover/internal/domain/size"
)

type Obstacle struct {
	Position *position.Position
}

func In(position *position.Position) *Obstacle {
	return &Obstacle{position}
}

func (this Obstacle) IsWithinLimit(size size.Size) bool {
	return this.Position.IsWithin(size)
}
