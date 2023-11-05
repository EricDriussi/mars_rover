package direction

import (
	"mars_rover/internal/domain/coordinate"
)

type West struct{}

func (this West) Degree() int {
	return 0
}

func (this West) DirectionOnTheLeft() Direction {
	return &South{}
}

func (this West) DirectionOnTheRight() Direction {
	return &North{}
}

func (this West) RelativePositionAhead() coordinate.RelativeCoordinate {
	return *coordinate.RelativeFrom(-1, 0)
}

func (this West) RelativePositionBehind() coordinate.RelativeCoordinate {
	return *coordinate.RelativeFrom(1, 0)
}
