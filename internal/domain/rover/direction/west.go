package direction

import (
	"mars_rover/internal/domain/coordinate/relativeCoordinate"
	. "mars_rover/internal/domain/coordinate/relativeCoordinate"
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
	return *relativeCoordinate.New(-1, 0)
}

func (this West) RelativeCoordinateBehind() RelativeCoordinate {
	return *relativeCoordinate.New(1, 0)
}