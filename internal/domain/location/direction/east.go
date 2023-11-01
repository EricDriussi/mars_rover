package direction

import (
	coordinate2d "mars_rover/internal/domain/coordinate/coordinate2D"
	relativePosition "mars_rover/internal/domain/location/relative_position"
)

type East struct{}

func (this East) CardinalPoint() string {
	return "E"
}

func (this East) DirectionOnTheLeft() Direction {
	return &North{}
}

func (this East) DirectionOnTheRight() Direction {
	return &South{}
}

func (this East) RelativePositionAhead() relativePosition.RelativePosition {
	return *relativePosition.New(coordinate2d.New(1, 0))
}

func (this East) RelativePositionBehind() relativePosition.RelativePosition {
	return *relativePosition.New(coordinate2d.New(-1, 0))
}
