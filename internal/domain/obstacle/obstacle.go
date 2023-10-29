package obstacle

import (
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/size"
)

type Obstacle struct {
	Location *location.Location
}

func In(location *location.Location) *Obstacle {
	return &Obstacle{location}
}

func (this Obstacle) IsWithinLimit(size size.Size) bool {
	return this.Location.IsWithin(size)
}
