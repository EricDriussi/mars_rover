package rover

import (
	. "github.com/google/uuid"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/rover/direction"
	. "mars_rover/internal/domain/rover/planetMap"
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
