package rover

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/direction"
	"mars_rover/internal/domain/planet"
	planetMap "mars_rover/internal/domain/planet_map"
)

type Rover struct {
	location    coordinate.Coordinate
	orientation direction.Direction
	planetMap   planetMap.Map
}

func Land(position coordinate.Coordinate, direction direction.Direction, planet planet.Planet) *Rover {
	return &Rover{location: position, orientation: direction, planetMap: *planetMap.Of(planet)}
}

func (this Rover) Direction() direction.Direction {
	return this.orientation
}

func (this Rover) Position() coordinate.Coordinate {
	return this.location
}

func (this *Rover) MoveForward() {
	this.location.AddOrWrap(this.orientation.RelativePositionAhead(), this.planetMap.Size())
}

func (this *Rover) MoveBackward() {
	this.location.AddOrWrap(this.orientation.RelativePositionBehind(), this.planetMap.Size())
}

func (this *Rover) TurnLeft() {
	this.orientation = this.orientation.DirectionOnTheLeft()
}

func (this *Rover) TurnRight() {
	this.orientation = this.orientation.DirectionOnTheRight()
}

func (this Rover) CheckObstacle() bool {
	return this.planetMap.CheckCollision(this.location)
}
