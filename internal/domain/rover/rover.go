package rover

import (
	"mars_rover/internal/domain/direction"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/planet"
	planetMap "mars_rover/internal/domain/planet_map"
)

type Rover struct {
	location    location.Location
	orientation direction.Direction
	planetMap   planetMap.Map
}

func Land(location location.Location, direction direction.Direction, planet planet.Planet) *Rover {
	return &Rover{location: location, orientation: direction, planetMap: *planetMap.Of(planet)}
}

func (this Rover) Direction() direction.Direction {
	return this.orientation
}

func (this Rover) Location() location.Location {
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
