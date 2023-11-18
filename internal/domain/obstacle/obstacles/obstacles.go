package obstacles

import (
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle"
	. "mars_rover/internal/domain/size"
)

type Obstacles struct {
	list []Obstacle
}

func New(list []Obstacle) *Obstacles {
	return &Obstacles{list}
}

func (this *Obstacles) List() []Obstacle {
	return this.list
}

func (this *Obstacles) Occupy(coordinate AbsoluteCoordinate) bool {
	for _, obs := range this.list {
		if obs.Occupies(coordinate) {
			return true
		}
	}
	return false
}

func (this *Obstacles) IsAnyBeyond(size Size) bool {
	for _, obs := range this.list {
		if obs.IsBeyond(size) {
			return true
		}
	}
	return false
}
