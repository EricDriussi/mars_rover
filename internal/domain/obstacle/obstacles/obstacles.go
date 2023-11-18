package obstacles

import (
	abs "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
)

type Obstacles struct {
	list []obstacle.Obstacle
}

func New(list []obstacle.Obstacle) *Obstacles {
	return &Obstacles{list}
}

func (this *Obstacles) List() []obstacle.Obstacle {
	return this.list
}

func (this *Obstacles) Occupy(coordinate abs.AbsoluteCoordinate) bool {
	for _, obs := range this.list {
		if obs.Occupies(coordinate) {
			return true
		}
	}
	return false
}

func (this *Obstacles) IsAnyBeyond(size size.Size) bool {
	for _, obs := range this.list {
		if obs.IsBeyond(size) {
			return true
		}
	}
	return false
}
