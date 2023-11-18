package rover

import (
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/direction"
	. "mars_rover/internal/domain/rover/planetMap"
)

type Rover interface {
	TurnLeft()
	TurnRight()
	MoveForward() error
	MoveBackward() error
	Position() AbsoluteCoordinate
	Direction() Direction
	Map() Map
}
