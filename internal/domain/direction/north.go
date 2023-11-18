package direction

import (
	"mars_rover/internal/domain/coordinate/relativeCoordinate"
	. "mars_rover/internal/domain/coordinate/relativeCoordinate"
)

type North struct{}

func (this North) CardinalPoint() string {
	return "N"
}

func (this North) DirectionOnTheLeft() Direction {
	return &West{}
}

func (this North) DirectionOnTheRight() Direction {
	return &East{}
}

func (this North) RelativePositionAhead() RelativeCoordinate {
	return *relativeCoordinate.New(0, 1)
}

func (this North) RelativePositionBehind() RelativeCoordinate {
	return *relativeCoordinate.New(0, -1)
}
