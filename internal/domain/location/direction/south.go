package direction

import (
	"mars_rover/internal/domain/coordinate"
)

type South struct{}

func (this South) CardinalPoint() string {
	return "S"
}

func (this South) DirectionOnTheLeft() Direction {
	return &East{}
}

func (this South) DirectionOnTheRight() Direction {
	return &West{}
}

func (this South) RelativePositionAhead() coordinate.RelativeCoordinate {
	return *coordinate.RelativeFrom(0, -1)
}

func (this South) RelativePositionBehind() coordinate.RelativeCoordinate {
	return *coordinate.RelativeFrom(0, 1)
}
