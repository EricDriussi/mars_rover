package relativePosition

import "mars_rover/internal/domain/coordinate"

type RelativePosition struct {
	coordinate *coordinate.Coordinate2D
}

func New(x, y int) *RelativePosition {
	return &RelativePosition{coordinate.New(x, y)}
}

func (this RelativePosition) X() int {
	return this.coordinate.X()
}

func (this RelativePosition) Y() int {
	return this.coordinate.Y()
}
