package direction

import (
	"mars_rover/src/domain/coordinate/relativeCoordinate"
	. "mars_rover/src/domain/coordinate/relativeCoordinate"
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
	return *relativeCoordinate.Down()
}

func (this South) RelativeCoordinateBehind() RelativeCoordinate {
	return *relativeCoordinate.Up()
}
