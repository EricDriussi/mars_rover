package rover

import (
	. "github.com/google/uuid"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
	. "mars_rover/src/domain/rover/planetMap"
)

type Rover interface {
	TurnLeft()
	TurnRight()
	MoveForward() error
	MoveBackward() error
	// TODO.LM: google's uuid should be wrapped
	Id() UUID
	Coordinate() AbsoluteCoordinate
	Direction() Direction
	Map() Map
}
