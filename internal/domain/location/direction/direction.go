package direction

import (
	"mars_rover/internal/domain/coordinate"
)

// TODO.LM: should this be a factory?
type Direction interface {
	Degree() int
	DirectionOnTheLeft() Direction
	DirectionOnTheRight() Direction
	RelativePositionAhead() coordinate.RelativeCoordinate
	RelativePositionBehind() coordinate.RelativeCoordinate
}
