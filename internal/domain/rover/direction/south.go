package direction

import (
	"mars_rover/internal/domain/coordinate/relativeCoordinate"
	. "mars_rover/internal/domain/coordinate/relativeCoordinate"
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

func (this South) RelativeCoordinateAhead() RelativeCoordinate {
	return *relativeCoordinate.New(0, -1)
}

func (this South) RelativeCoordinateBehind() RelativeCoordinate {
	return *relativeCoordinate.New(0, 1)
}
