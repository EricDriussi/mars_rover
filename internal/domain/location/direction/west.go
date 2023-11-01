package direction

import (
	coordinate2d "mars_rover/internal/domain/coordinate/coordinate2D"
	relativePosition "mars_rover/internal/domain/location/relative_position"
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
	return *relativePosition.New(coordinate2d.New(-1, 0))
}

func (this West) RelativePositionBehind() relativePosition.RelativePosition {
	return *relativePosition.New(coordinate2d.New(1, 0))
}
