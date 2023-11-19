package direction

import (
	. "mars_rover/internal/domain/coordinate/relativeCoordinate"
)

// TODO: move package inside rover package
type Direction interface {
	CardinalPoint() string
	DirectionOnTheLeft() Direction
	DirectionOnTheRight() Direction
	RelativeCoordinateAhead() RelativeCoordinate
	RelativeCoordinateBehind() RelativeCoordinate
}
