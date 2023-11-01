package direction

import (
	coordinate2d "mars_rover/internal/domain/coordinate/coordinate2D"
	relativePosition "mars_rover/internal/domain/location/relative_position"
)

type South struct{}

func (this South) CardinalPoint() string {
	return "S"
}

func (this South) DirectionOnTheLeft() Direction {
	return &East{}
}

func (this South) DirectionOnTheRight() Direction {
	return &West{}
}

func (this South) RelativePositionAhead() relativePosition.RelativePosition {
	return *relativePosition.New(coordinate2d.New(0, -1))
}

func (this South) RelativePositionBehind() relativePosition.RelativePosition {
	return *relativePosition.New(coordinate2d.New(0, 1))
}
