package direction

import (
	"mars_rover/internal/domain/coordinate"
)

type North struct{}

func (this North) Degree() int {
	return 90
}

func (this North) DirectionOnTheLeft() Direction {
	return &West{}
}

func (this North) DirectionOnTheRight() Direction {
	return &East{}
}

func (this North) RelativePositionAhead() coordinate.RelativeCoordinate {
	return *coordinate.RelativeFrom(0, 1)
}

func (this North) RelativePositionBehind() coordinate.RelativeCoordinate {
	return *coordinate.RelativeFrom(0, -1)
}
