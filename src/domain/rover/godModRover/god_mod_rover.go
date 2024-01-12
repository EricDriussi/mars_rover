package godModRover

import (
	"mars_rover/src/domain/coordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover/direction"
	. "mars_rover/src/domain/rover/id"
	"mars_rover/src/domain/rover/planetMap"
	. "mars_rover/src/domain/rover/planetMap"
)

// TODO.LM: This Rover doesn't care about collisions or size limits
type GodModRover struct {
	id         ID
	planetMap  Map
	coordinate AbsoluteCoordinate
	direction  Direction
}

func LandFacing(id ID, direction Direction, coordinate AbsoluteCoordinate, planet Planet) *GodModRover {
	return &GodModRover{
		id:         id,
		planetMap:  *planetMap.OfPlanet(planet),
		coordinate: coordinate,
		direction:  direction,
	}
}

func (this *GodModRover) MoveForward() error {
	this.coordinate = *coordinate.SumOf(this.coordinate, this.direction.RelativeCoordinateAhead())
	return nil
}

func (this *GodModRover) MoveBackward() error {
	this.coordinate = *coordinate.SumOf(this.coordinate, this.direction.RelativeCoordinateBehind())
	return nil
}

func (this *GodModRover) TurnLeft() {
	this.direction = this.direction.DirectionOnTheLeft()
}

func (this *GodModRover) TurnRight() {
	this.direction = this.direction.DirectionOnTheRight()
}

func (this *GodModRover) Id() ID {
	return this.id
}

func (this *GodModRover) Coordinate() AbsoluteCoordinate {
	return this.coordinate
}

func (this *GodModRover) Direction() Direction {
	return this.direction
}

func (this *GodModRover) Map() Map {
	return this.planetMap
}
