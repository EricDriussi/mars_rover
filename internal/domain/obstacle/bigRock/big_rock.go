package bigRock

import (
	coord "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
)

type BigRock struct {
	coordinates []coord.AbsoluteCoordinate
}

func In(coordinate []coord.AbsoluteCoordinate) obstacle.Obstacle {
	return &BigRock{coordinate}
}

func (this *BigRock) Occupies(coordinate coord.AbsoluteCoordinate) bool {
	for _, occupiedCoordinate := range this.coordinates {
		if coordinate.Equals(&occupiedCoordinate) {
			return true
		}
	}
	return false
}

func (this *BigRock) IsBeyond(size size.Size) bool {
	for _, occupiedCoordinate := range this.coordinates {
		if occupiedCoordinate.X() > size.Width() || occupiedCoordinate.Y() > size.Height() {
			return true
		}
	}
	return false
}

func (this *BigRock) Coordinates() []coord.AbsoluteCoordinate {
	return this.coordinates
}
