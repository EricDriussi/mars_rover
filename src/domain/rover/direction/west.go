package direction

import (
	"mars_rover/src/domain/coordinate/relativeCoordinate"
	. "mars_rover/src/domain/coordinate/relativeCoordinate"
)

type West struct{}

func (this West) CardinalPoint() string {
	return "W"
}

func (this West) DirectionOnTheLeft() Direction {
	return &South{}
}

func (this West) DirectionOnTheRight() Direction {
	return &North{}
}

func (this West) RelativeCoordinateAhead() RelativeCoordinate {
	return *relativeCoordinate.Left()
}

func (this West) RelativeCoordinateBehind() RelativeCoordinate {
	return *relativeCoordinate.Right()
}
