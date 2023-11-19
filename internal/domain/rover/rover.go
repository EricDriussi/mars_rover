package rover

import (
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/rover/direction"
	. "mars_rover/internal/domain/rover/planetMap"
)

type Rover interface {
	TurnLeft()
	TurnRight()
	MoveForward() error
	MoveBackward() error
	Coordinate() AbsoluteCoordinate
	Direction() Direction
	Map() Map
}
