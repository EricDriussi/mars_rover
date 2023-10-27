package direction

import (
	relativePosition "mars_rover/internal/domain/relative_position"
)

type West struct{}

func (this West) CardinalPoint() string {
	return "W"
}

func (this West) DirectionOnTheLeft() Direction {
	return &South{}
}

func (this West) DirectionOnTheRight() Direction {
	return &North{}
}

func (this West) RelativePositionAhead() relativePosition.RelativePosition {
	return *relativePosition.New(-1, 0)
}

func (this West) RelativePositionBehind() relativePosition.RelativePosition {
	return *relativePosition.New(1, 0)
}
