package direction

import (
	"mars_rover/internal/domain/coordinate"
)

type Direction interface {
	CardinalPoint() string
	DirectionOnTheLeft() Direction
	DirectionOnTheRight() Direction
	RelativePositionAhead() coordinate.RelativeCoordinate
	RelativePositionBehind() coordinate.RelativeCoordinate
}
