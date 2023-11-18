package direction

import (
	"mars_rover/internal/domain/coordinate/relativeCoordinate"
	. "mars_rover/internal/domain/coordinate/relativeCoordinate"
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

func (this East) RelativePositionAhead() RelativeCoordinate {
	return *relativeCoordinate.New(1, 0)
}

func (this East) RelativePositionBehind() RelativeCoordinate {
	return *relativeCoordinate.New(-1, 0)
}
