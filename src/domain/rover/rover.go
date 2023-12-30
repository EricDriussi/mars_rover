package rover

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
	. "mars_rover/src/domain/rover/planetMap"
	. "mars_rover/src/domain/rover/uuid"
)

type Rover interface {
	TurnLeft()
	TurnRight()
	MoveForward() error
	MoveBackward() error
	Id() UUID
	Coordinate() AbsoluteCoordinate
	Direction() Direction
	Map() Map
}
