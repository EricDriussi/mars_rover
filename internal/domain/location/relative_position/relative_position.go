package relativePosition

import "mars_rover/internal/domain/coordinate"

type RelativePosition struct {
	coordinate coordinate.Coordinate
}

func New(coord coordinate.Coordinate) *RelativePosition {
	return &RelativePosition{coord}
}

func (this RelativePosition) X() int {
	return this.coordinate.X()
}

func (this RelativePosition) Y() int {
	return this.coordinate.Y()
}
