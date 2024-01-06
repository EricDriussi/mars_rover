package obstacles

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	. "mars_rover/src/domain/size"
)

type Obstacles struct {
	list []Obstacle
}

func FromList(list ...Obstacle) *Obstacles {
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

func (this *Obstacles) Amount() int {
	return len(this.list)
}

func (this *Obstacles) Add(obstacle Obstacle) {
	this.list = append(this.list, obstacle)
}
