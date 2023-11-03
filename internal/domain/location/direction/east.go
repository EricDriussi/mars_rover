package direction

import (
	"mars_rover/internal/domain/coordinate"
)

type East struct{}

func (this East) CardinalPoint() string {
	return "E"
}

func (this East) DirectionOnTheLeft() Direction {
	return &North{}
}

func (this East) DirectionOnTheRight() Direction {
	return &South{}
}

func (this East) RelativePositionAhead() coordinate.RelativeCoordinate {
	return *coordinate.RelativeFrom(1, 0)
}

func (this East) RelativePositionBehind() coordinate.RelativeCoordinate {
	return *coordinate.RelativeFrom(-1, 0)
}
