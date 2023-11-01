package direction

import (
	relativePosition "mars_rover/internal/domain/location/relative_position"
)

type North struct{}

func (this North) CardinalPoint() string {
	return "N"
}

func (this North) DirectionOnTheLeft() Direction {
	return &West{}
}

func (this North) DirectionOnTheRight() Direction {
	return &East{}
}

func (this North) RelativePositionAhead() relativePosition.RelativePosition {
	return *relativePosition.New(0, 1)
}

func (this North) RelativePositionBehind() relativePosition.RelativePosition {
	return *relativePosition.New(0, -1)
}
