package direction

import (
	"mars_rover/src/domain/coordinate/relativeCoordinate"
	. "mars_rover/src/domain/coordinate/relativeCoordinate"
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

func (this East) RelativeCoordinateAhead() RelativeCoordinate {
	return *relativeCoordinate.Right()
}

func (this East) RelativeCoordinateBehind() RelativeCoordinate {
	return *relativeCoordinate.Left()
}
