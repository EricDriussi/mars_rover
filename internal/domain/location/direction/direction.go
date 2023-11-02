package direction

import (
	relativePosition "mars_rover/internal/domain/location/relative_position"
)

type Direction interface {
	// TODO: should be renamed to Degree and return a 0-360 int
	CardinalPoint() string
	DirectionOnTheLeft() Direction
	DirectionOnTheRight() Direction
	RelativePositionAhead() relativePosition.RelativePosition
	RelativePositionBehind() relativePosition.RelativePosition
}
