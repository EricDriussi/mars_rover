package direction

import (
	relativePosition "mars_rover/internal/domain/relative_position"
)

type Direction interface {
	CardinalPoint() string
	DirectionOnTheLeft() Direction
	DirectionOnTheRight() Direction
	RelativePositionAhead() relativePosition.RelativePosition
	RelativePositionBehind() relativePosition.RelativePosition
}
