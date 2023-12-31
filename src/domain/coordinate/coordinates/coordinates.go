package coordinates

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/size"
)

type Coordinates struct {
	list []AbsoluteCoordinate
}

func New(list ...AbsoluteCoordinate) *Coordinates {
	return &Coordinates{filterUnique(list)}
}

func filterUnique(list []AbsoluteCoordinate) []AbsoluteCoordinate {
	uniqueList := make([]AbsoluteCoordinate, 0)
	for _, coordinate := range list {
		if !contains(uniqueList, coordinate) {
			uniqueList = append(uniqueList, coordinate)
		}
	}
	return uniqueList
}

func (this *Coordinates) Contain(coordinate AbsoluteCoordinate) bool {
	return contains(this.list, coordinate)
}

func contains(list []AbsoluteCoordinate, coordinate AbsoluteCoordinate) bool {
	for _, containedCoordinate := range list {
		if containedCoordinate.Equals(coordinate) {
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
