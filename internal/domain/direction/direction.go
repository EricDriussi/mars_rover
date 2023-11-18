package direction

import (
	. "mars_rover/internal/domain/coordinate/relativeCoordinate"
)

type Direction interface {
	CardinalPoint() string
	DirectionOnTheLeft() Direction
	DirectionOnTheRight() Direction
	RelativeCoordinateAhead() RelativeCoordinate
	RelativeCoordinateBehind() RelativeCoordinate
}
