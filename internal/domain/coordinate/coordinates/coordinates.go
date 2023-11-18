package coordinates

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/size"
)

type Coordinates struct {
	list []absoluteCoordinate.AbsoluteCoordinate
}

func New(coordinate []absoluteCoordinate.AbsoluteCoordinate) *Coordinates {
	return &Coordinates{coordinate}
}

func (this *Coordinates) Contain(coordinate absoluteCoordinate.AbsoluteCoordinate) bool {
	for _, occupiedCoordinate := range this.list {
		if coordinate.Equals(&occupiedCoordinate) {
			return true
		}
	}
	return false
}

func (this *Coordinates) GoBeyond(size size.Size) bool {
	for _, occupiedCoordinate := range this.list {
		if occupiedCoordinate.X() > size.Width() || occupiedCoordinate.Y() > size.Height() {
			return true
		}
	}
	return false
}

func (this *Coordinates) List() []absoluteCoordinate.AbsoluteCoordinate {
	return this.list
}
