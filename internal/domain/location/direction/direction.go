package direction

import (
	"mars_rover/internal/domain/coordinate"
)

type Direction interface {
	// TODO: should be renamed to Degree and return a 0-360 int
	CardinalPoint() string
	DirectionOnTheLeft() Direction
	DirectionOnTheRight() Direction
	RelativePositionAhead() coordinate.RelativeCoordinate
	RelativePositionBehind() coordinate.RelativeCoordinate
}
