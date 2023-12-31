package coordinates

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/size"
)

type Coordinates struct {
	list []AbsoluteCoordinate
}

func New(list ...AbsoluteCoordinate) *Coordinates {
	return &Coordinates{list}
}

func (this *Coordinates) Contain(coordinate AbsoluteCoordinate) bool {
	for _, containedCoordinate := range this.list {
		if coordinate.Equals(containedCoordinate) {
			return true
		}
	}
	return false
}

func (this *Coordinates) Overflow(size Size) bool {
	for _, coordinate := range this.list {
		if coordinate.X() > size.Width() || coordinate.Y() > size.Height() {
			return true
		}
	}
	return false
}

func (this *Coordinates) List() []AbsoluteCoordinate {
	return this.list
}
