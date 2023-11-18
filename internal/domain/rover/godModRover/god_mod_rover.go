package godModRover

import (
	"mars_rover/internal/domain/coordinate"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/direction"
	. "mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover/planetMap"
	. "mars_rover/internal/domain/rover/planetMap"
)

type GodModRover struct {
	planetMap  Map
	coordinate AbsoluteCoordinate
	direction  Direction
}

func Land(coordinate AbsoluteCoordinate, planet Planet) *GodModRover {
	return &GodModRover{
		planetMap:  *planetMap.Of(planet),
		coordinate: coordinate,
		direction:  North{},
	}
}

func LandFacing(direction Direction, coordinate AbsoluteCoordinate, planet Planet) *GodModRover {
	return &GodModRover{
		planetMap:  *planetMap.Of(planet),
		coordinate: coordinate,
		direction:  direction,
	}
}

func (this *GodModRover) TurnLeft() {
	this.direction = this.direction.DirectionOnTheLeft()
}

func (this *GodModRover) TurnRight() {
	this.direction = this.direction.DirectionOnTheRight()
}

func (this *GodModRover) MoveForward() error {
	this.coordinate = *coordinate.SumOf(this.coordinate, this.direction.RelativePositionAhead())
	return nil
}

func (this *GodModRover) MoveBackward() error {
	this.coordinate = *coordinate.SumOf(this.coordinate, this.direction.RelativePositionBehind())
	return nil
}

func (this *GodModRover) Position() AbsoluteCoordinate {
	return this.coordinate
}

func (this *GodModRover) Direction() Direction {
	return this.direction
}

func (this *GodModRover) Map() Map {
	return this.planetMap
}
