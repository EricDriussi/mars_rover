package rover

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
	. "mars_rover/src/domain/rover/id"
	. "mars_rover/src/domain/rover/planetMap"
)

type Rover interface {
	TurnLeft()
	TurnRight()
	MoveForward() error
	MoveBackward() error
	Id() ID
	Coordinate() AbsoluteCoordinate
	Direction() Direction
	Map() Map
}
